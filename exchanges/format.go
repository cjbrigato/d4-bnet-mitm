package exchanges

import (
	"encoding/hex"
	"fmt"

	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/notification/v2/client"
	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ResolveNotificationPayload(bgs_rpc_message_bytes []byte) {
	notifmsg := &client.NotificationReceivedNotification{}
	if err := proto.Unmarshal(bgs_rpc_message_bytes, notifmsg); err != nil {
		fmt.Printf("Failed to parse Notification: %s\n", err)
	}
	notif := notifmsg.GetNotifications()[0].GetAttribute()
	messageid := notif[0].GetValue().GetIntValue()
	payload := notif[1].GetValue().GetBlobValue()
	var messageid_type string
	switch messageid {
	case 0:
		messageid_type = "Fenris.ClientMessage.FindUserProxyResponse"
	case 5:
		messageid_type = "Fenris.ClientMessage.PingConnectInfoSingleResult"
	case 2:
		messageid_type = "Fenris.ClientMessage.QueueUpdate"
	default:
		messageid_type = "unknown"
	}
	fmt.Printf("### %s (as FEN.NotificationMessage.Payload)\n", messageid_type)
	if messageid_type != "unknown" {
		FenrisMessage, err := dynamic.ParseAs(messageid_type, payload)
		if err != nil {
			fmt.Printf("Failed to parse %s: %s\n", messageid_type, err)
		} else {
			PrintMessage(*FenrisMessage)
		}
	} else {
		fmt.Printf("%s\n", hex.Dump(payload))
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
			// Finish printing the value.
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
