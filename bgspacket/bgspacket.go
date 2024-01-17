package bgspacket

import (
	"embed"
	"encoding/binary"
	"fmt"

	"github.com/cjbrigato/d4-bnet-mitm/bnet/Fenris/ClientMessage"
	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/notification/v2/client"
	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	"github.com/cjbrigato/d4-bnet-mitm/log"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtocolInit(checkRegistry bool, f *embed.FS) {
	dynamic.Register("vfs/bgs_bundle.binpb", f)
	dynamic.Register("vfs/fenris_bundle.binpb", f)
	if checkRegistry {
		services.Test_protos()
	}
}

type BgsHeader struct {
	Bytes       []byte
	Len         uint16
	ProtoHeader *protocol.Header
}

func (h BgsHeader) PrintMessage() {
	PrintMessage(h.ProtoHeader)
}

type BgsMessage struct {
	Bytes        []byte
	ProtoMessage proto.Message
}

func (m BgsMessage) PrintMessage() {
	PrintMessage(m.ProtoMessage)
}

type BgsPacket struct {
	Frame           *ws.Frame
	BGSHeader       *BgsHeader
	BGSMessage      *BgsMessage
	ServiceHash     uint32
	MethodID        uint32
	RpcToken        uint32
	ServiceID       uint32
	rpc_kind        string
	MessageType     protoreflect.FullName
	FenNotification *FenNotification
}

func (p *BgsPacket) WSEncapsulate(frame ws.Frame) []byte {
	bgs_rpc_message_len := len(p.BGSMessage.Bytes)
	header_len := uint16(len(p.BGSHeader.Bytes))
	header_len_bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(header_len_bytes, header_len)
	newframepayload := make([]byte, len(header_len_bytes)+len(p.BGSHeader.Bytes)+len(p.BGSMessage.Bytes))
	newframepayload = append(header_len_bytes, p.BGSHeader.Bytes...)
	if bgs_rpc_message_len > 0 {
		newframepayload = append(newframepayload, p.BGSMessage.Bytes...)
	}
	return append(frame.Header.Bytes, newframepayload...)
}

type FenNotification struct {
	MessageBytes                   []byte
	ProtoMessage                   *client.NotificationReceivedNotification
	MessageID                      int64
	MessagePayloadBytes            []byte
	MessagePayloadProtoMessage     proto.Message
	MessagePayloadProtoMessageType protoreflect.FullName
}

type PacketSource int64

//go:generate stringer -type=PacketSource
const (
	SRC_CLIENT PacketSource = iota
	SRC_SERVER
	SRC_UNKNOWN
)

type PacketEntry struct {
	Source PacketSource
	Packet *BgsPacket
}

var PacketHistory = make(map[int64]PacketEntry)
var PacketHistoryCounter int64

func NewBgsPacketFromFrame(frame *ws.Frame, Source PacketSource, shouldCraftResponse bool) *BgsPacket {
	bgs_packet := &BgsPacket{Frame: frame}
	bgs_packet.BGSHeader = &BgsHeader{}
	bgs_packet.BGSMessage = &BgsMessage{}
	bgs_packet.BGSHeader.Len = binary.BigEndian.Uint16(frame.Payload[0:])
	bgs_packet.BGSHeader.Bytes = frame.Payload[2 : bgs_packet.BGSHeader.Len+2]
	bgs_rpc_message_len := len(frame.Payload) - int(bgs_packet.BGSHeader.Len+2)
	if bgs_rpc_message_len > 0 {
		bgs_packet.BGSMessage.Bytes = frame.Payload[2+bgs_packet.BGSHeader.Len:]
	}
	bgs_packet.BGSHeader.ProtoHeader = &protocol.Header{}
	if err := proto.Unmarshal(bgs_packet.BGSHeader.Bytes, bgs_packet.BGSHeader.ProtoHeader); err != nil {
		log.Error(nil, "Failed to parse BgsPacket Header: %s", err)
	}
	bgs_packet.ServiceHash = 0
	bgs_packet.MethodID = 0
	bgs_packet.ServiceID = bgs_packet.BGSHeader.ProtoHeader.GetServiceId()
	bgs_packet.RpcToken = bgs_packet.BGSHeader.ProtoHeader.GetToken()
	bgs_packet.rpc_kind = services.RPCCallKind(bgs_packet.ServiceID)
	switch bgs_packet.rpc_kind {
	case "request":
		bgs_packet.ServiceHash = bgs_packet.BGSHeader.ProtoHeader.GetServiceHash()
		bgs_packet.MethodID = bgs_packet.BGSHeader.ProtoHeader.GetMethodId()
		add_pending_response(TrackingToken(bgs_packet.RpcToken), PendingResponse{bgs_packet.ServiceHash, bgs_packet.MethodID})
	case "response":
		pending, ok := recall_pending_response(TrackingToken(bgs_packet.RpcToken))
		if ok {
			bgs_packet.ServiceHash = pending.ServiceHash
			bgs_packet.MethodID = pending.MethodId
		} else {
			log.Warn(&map[string]any{"NODECODE": "Cannot decode message Service or method"}, "Failure recalling a PendingResponse for Token %d", bgs_packet.RpcToken)
		}
	}
	if bgs_packet.ServiceHash > 0 {
		val, ok := services.Get(bgs_packet.ServiceHash)
		errorMap := make(map[string]any)
		if ok {
			traceMap := make(map[string]any)
			traceMap["service_hash"] = bgs_packet.ServiceHash
			traceMap["service_name"] = val.Name()
			traceMap["method_id"] = bgs_packet.MethodID
			if bgs_packet.MethodID > 0 {
				mval, mok := val.Method(uint16(bgs_packet.MethodID))
				if mok {
					traceMap["method_name"] = mval
					messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, bgs_packet.ServiceID))
					bgs_packet.MessageType = messageName
					traceMap["method_name"] = messageName
				} else {
					errorMap["method_name"] = fmt.Sprintf("Unknown method: %s", mval)
					log.Error(&errorMap, "Dynamic Protobuf MessageType computation FAILURE")
				}
			}
			log.Trace(&traceMap, "Dynamic Protobuf MessageType computation infos:")
		} else {
			errorMap["service_hash"] = fmt.Sprintf("Unknown Service hash: %d", bgs_packet.ServiceHash)
			log.Error(&errorMap, "Dynamic Protobuf MessageType computation FAILURE")
		}
	}
	if len(bgs_packet.BGSMessage.Bytes) > 0 {
		msg, err := services.DecodeServiceMessageFromType(bgs_packet.MessageType, bgs_packet.BGSMessage.Bytes)
		if err != nil {
			log.Error(nil, fmt.Sprintf("%v", err))
		} else {
			bgs_packet.BGSMessage.ProtoMessage = *msg
			if fmt.Sprintf("%s", bgs_packet.MessageType) == "bgs.protocol.notification.v2.client.NotificationReceivedNotification" {
				bgs_packet.resolveNotificationPayload(shouldCraftResponse)
			}
		}

	}
	PacketHistory[PacketHistoryCounter] = PacketEntry{Source: Source, Packet: bgs_packet}
	PacketHistoryCounter++
	return bgs_packet
}
func (p *BgsPacket) resolveNotificationPayload(shouldCraftResponse bool) {
	p.FenNotification = &FenNotification{}
	p.FenNotification.MessagePayloadProtoMessage = proto.Message.ProtoReflect(&ClientMessage.FindUserProxyResponse{}).Interface()
	p.FenNotification.ProtoMessage = &client.NotificationReceivedNotification{}
	p.FenNotification.MessageBytes = make([]byte, len(p.BGSMessage.Bytes))
	copy(p.FenNotification.MessageBytes, p.BGSMessage.Bytes)
	log.Trace(nil, "Adding / Solving Attached FENNotification")
	if err := proto.Unmarshal(p.FenNotification.MessageBytes, p.FenNotification.ProtoMessage); err != nil {
		log.Error(nil, "Failed to parse Notification: %s", err)
		return
	}
	notif := p.FenNotification.ProtoMessage.GetNotifications()[0].GetAttribute()
	p.FenNotification.MessageID = notif[0].GetValue().GetIntValue()
	p.FenNotification.MessagePayloadBytes = notif[1].GetValue().GetBlobValue()
	switch p.FenNotification.MessageID {
	case 0:
		p.FenNotification.MessagePayloadProtoMessageType = "Fenris.ClientMessage.FindUserProxyResponse"
		p.FenNotification.MessagePayloadProtoMessage = proto.Message.ProtoReflect(&ClientMessage.FindUserProxyResponse{}).Interface()
		if err := proto.Unmarshal(p.FenNotification.MessagePayloadBytes, p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse)); err != nil {
			log.Error(nil, "Failed to parse FindUserProxyResponse: %s", err)
		}
		if shouldCraftResponse {
			log.Info(nil, "Crafting a response for %s", p.FenNotification.MessagePayloadProtoMessageType)
			var err error
			p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse).ConnectInfo.Address = proto.String("127.0.0.1:61000")
			p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse).ConnectInfo.Port = proto.Uint32(61000)
			p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse).ConnectInfo.Player[0].EncryptionInfo.PresharedKey.Modulus = []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
			p.FenNotification.ProtoMessage.GetNotifications()[0].GetAttribute()[1].Value.BlobValue, err = proto.Marshal(p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse))
			if err != nil {
				log.Error(nil, "Failed to Marshal crafted FindUserProxyResponse: %s", err)
			}
			p.BGSMessage.Bytes, err = proto.Marshal(p.FenNotification.ProtoMessage)
			if err != nil {
				log.Error(nil, "Failed to Marshal crafted : %s", err)
			}
			*p.BGSHeader.ProtoHeader.Size = uint32(len(p.BGSMessage.Bytes))
			p.BGSHeader.Bytes, err = proto.Marshal(p.BGSHeader.ProtoHeader)
		}
	case 5:
		p.FenNotification.MessagePayloadProtoMessageType = "Fenris.ClientMessage.PingConnectInfoSingleResult"
		p.FenNotification.MessagePayloadProtoMessage = proto.Message.ProtoReflect(&ClientMessage.PingConnectInfoSingleResult{}).Interface()
		if err := proto.Unmarshal(p.FenNotification.MessagePayloadBytes, p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.PingConnectInfoSingleResult)); err != nil {
			log.Error(nil, "Failed to parse FindUserProxyResponse: %s", err)
		}
	case 2:
		p.FenNotification.MessagePayloadProtoMessageType = "Fenris.ClientMessage.QueueUpdate"
		p.FenNotification.MessagePayloadProtoMessage = proto.Message.ProtoReflect(&ClientMessage.QueueUpdate{}).Interface()
		if err := proto.Unmarshal(p.FenNotification.MessagePayloadBytes, p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.QueueUpdate)); err != nil {
			log.Error(nil, "Failed to parse FindUserProxyResponse: %s", err)
		}
	default:
		//var err error
		p.FenNotification.MessagePayloadProtoMessageType = "unknown"
		/*pbMessage := &p.FenNotification.MessagePayloadProtoMessage
		pbMessage, err = dynamic.ParseAs(fmt.Sprintf("%s", p.FenNotification.MessagePayloadProtoMessageType), p.FenNotification.MessagePayloadBytes)
		if err != nil {
			log.Error(nil, "Failed to parse %s: %s", messageid_type, err)
		} else {
			p.FenNotification.MessagePayloadProtoMessage = *pbMessage
		}*/
	}
}

func PrintMessage(m proto.Message) {
	var indent []byte
	protorange.Options{
		Stable: true,
	}.Range(m.ProtoReflect(),
		func(p protopath.Values) error {
			// Print the key.
			var fd protoreflect.FieldDescriptor
			last := p.Index(-1)
			beforeLast := p.Index(-2)
			switch last.Step.Kind() {
			case protopath.FieldAccessStep:
				fd = last.Step.FieldDescriptor()
				fmt.Printf("%s%s: ", indent, fd.Name())
			case protopath.ListIndexStep:
				fd = beforeLast.Step.FieldDescriptor() // lists always appear in the context of a repeated field
				fmt.Printf("%s%d: ", indent, last.Step.ListIndex())
			case protopath.MapIndexStep:
				fd = beforeLast.Step.FieldDescriptor() // maps always appear in the context of a repeated field
				fmt.Printf("%s%v: ", indent, last.Step.MapIndex().Interface())
			case protopath.AnyExpandStep:
				fmt.Printf("%s[%v]: ", indent, last.Value.Message().Descriptor().FullName())
			case protopath.UnknownAccessStep:
				fmt.Printf("%s?: ", indent)
			}

			// Starting printing the value.
			switch v := last.Value.Interface().(type) {
			case protoreflect.Message:
				fmt.Printf("{\n")
				indent = append(indent, ' ', ' ')
			case protoreflect.List:
				fmt.Printf("[\n")
				indent = append(indent, ' ', ' ')
			case protoreflect.Map:
				fmt.Printf("{\n")
				indent = append(indent, ' ', ' ')
			case protoreflect.EnumNumber:
				var ev protoreflect.EnumValueDescriptor
				if fd != nil {
					ev = fd.Enum().Values().ByNumber(v)
				}
				if ev != nil {
					fmt.Printf("%v\n", ev.Name())
				} else {
					fmt.Printf("%v\n", v)
				}
			case string, []byte:
				fmt.Printf("%q\n", v)
			default:
				fmt.Printf("%v\n", v)
			}
			return nil
		},
		func(p protopath.Values) error {
			last := p.Index(-1)
			switch last.Value.Interface().(type) {
			case protoreflect.Message:
				indent = indent[:len(indent)-2]
				fmt.Printf("%s}\n", indent)
			case protoreflect.List:
				indent = indent[:len(indent)-2]
				fmt.Printf("%s]\n", indent)
			case protoreflect.Map:
				indent = indent[:len(indent)-2]
				fmt.Printf("%s}\n", indent)
			}
			return nil
		},
	)
}

func truncate(str string, length int) (truncated string, was_truncated bool) {
	was_truncated = false
	if length <= 0 {
		return
	}
	for i, char := range str {
		if i >= length {
			was_truncated = true
			break
		}
		truncated += string(char)
	}
	return
}
