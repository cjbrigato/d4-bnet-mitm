package bgspacket

import (
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
	"google.golang.org/protobuf/reflect/protoreflect"
)

type BgsHeader struct {
	Bytes        []byte
	Len          uint16
	proto_header *protocol.Header
}

type BgsMessage struct {
	Bytes         []byte
	proto_message proto.Message
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

func NewBgsPacketFromFrame(frame *ws.Frame, shouldCraftResponse bool) *BgsPacket {
	bgs_packet := &BgsPacket{Frame: frame}
	bgs_packet.BGSHeader = &BgsHeader{}
	bgs_packet.BGSMessage = &BgsMessage{}
	bgs_packet.BGSHeader.Len = binary.BigEndian.Uint16(frame.Payload[0:])
	bgs_packet.BGSHeader.Bytes = frame.Payload[2 : bgs_packet.BGSHeader.Len+2]
	bgs_rpc_message_len := len(frame.Payload) - int(bgs_packet.BGSHeader.Len+2)
	if bgs_rpc_message_len > 0 {
		bgs_packet.BGSMessage.Bytes = frame.Payload[2+bgs_packet.BGSHeader.Len:]
	}
	bgs_packet.BGSHeader.proto_header = &protocol.Header{}
	if err := proto.Unmarshal(bgs_packet.BGSHeader.Bytes, bgs_packet.BGSHeader.proto_header); err != nil {
		log.Error(nil, "Failed to parse BgsPacket Header: %s", err)
	}
	bgs_packet.ServiceHash = 0
	bgs_packet.MethodID = 0
	bgs_packet.ServiceID = bgs_packet.BGSHeader.proto_header.GetServiceId()
	bgs_packet.RpcToken = bgs_packet.BGSHeader.proto_header.GetToken()
	bgs_packet.rpc_kind = services.RPCCallKind(bgs_packet.ServiceID)
	switch bgs_packet.rpc_kind {
	case "request":
		bgs_packet.ServiceHash = bgs_packet.BGSHeader.proto_header.GetServiceHash()
		bgs_packet.MethodID = bgs_packet.BGSHeader.proto_header.GetMethodId()
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
			bgs_packet.BGSMessage.proto_message = *msg
			if fmt.Sprintf("%s", bgs_packet.MessageType) == "bgs.protocol.notification.v2.client.NotificationReceivedNotification" {
				bgs_packet.resolveNotificationPayload(shouldCraftResponse)
			}
		}

	}
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
	var messageid_type string
	switch p.FenNotification.MessageID {
	case 0:
		p.FenNotification.MessagePayloadProtoMessageType = "Fenris.ClientMessage.FindUserProxyResponse"
		p.FenNotification.MessagePayloadProtoMessage = proto.Message.ProtoReflect(&ClientMessage.FindUserProxyResponse{}).Interface()
		if err := proto.Unmarshal(p.FenNotification.MessagePayloadBytes, p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse)); err != nil {
			log.Error(nil, "Failed to parse FindUserProxyResponse: %s", err)
		}
		if shouldCraftResponse {
			var err error
			p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse).ConnectInfo.Address = proto.String("127.0.0.1:61000")
			p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse).ConnectInfo.Port = proto.Uint32(61000)
			p.FenNotification.ProtoMessage.GetNotifications()[0].GetAttribute()[1].Value.BlobValue, err = proto.Marshal(p.FenNotification.MessagePayloadProtoMessage.(*ClientMessage.FindUserProxyResponse))
			if err != nil {
				log.Error(nil, "Failed to Marshal crafted FindUserProxyResponse: %s", err)
			}
			p.BGSMessage.Bytes, err = proto.Marshal(p.FenNotification.ProtoMessage)
			if err != nil {
				log.Error(nil, "Failed to Marshal crafted : %s", err)
			}
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
		var err error
		p.FenNotification.MessagePayloadProtoMessageType = "unknown"
		pbMessage := &p.FenNotification.MessagePayloadProtoMessage
		pbMessage, err = dynamic.ParseAs(fmt.Sprintf("%s", p.FenNotification.MessagePayloadProtoMessageType), p.FenNotification.MessagePayloadBytes)
		if err != nil {
			log.Error(nil, "Failed to parse %s: %s", messageid_type, err)
		} else {
			p.FenNotification.MessagePayloadProtoMessage = *pbMessage
		}
	}
}
