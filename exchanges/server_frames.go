package exchanges

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
	"github.com/gookit/color"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func HandleServer(r io.Reader, c io.Writer, source string, conn_id int64) {

	upgraded := false
	data := make([]byte, 512)
	for {
		if upgraded {
			frame, err := ws.ReadFrame(r)
			if err != nil {
				color.Danger.Printf("unable to read frame: %v", err)
				break
			}

			LastReadFrame := []byte{}
			LastReadFrame = append(LastReadFrame, frame.Header.Bytes...)
			LastReadFrame = append(LastReadFrame, frame.Payload...)

			n, err := c.Write(LastReadFrame)
			color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)

			Log_Mutex.Lock()
			bgs_rpc_header_len := binary.BigEndian.Uint16(frame.Payload[0:])
			bgs_rpc_header_bytes := frame.Payload[2 : bgs_rpc_header_len+2]
			bgs_rpc_message_len := len(frame.Payload) - int(bgs_rpc_header_len+2)
			bgs_rpc_message_bytes := []byte{'0'}
			if bgs_rpc_message_len > 0 {
				bgs_rpc_message_bytes = frame.Payload[2+bgs_rpc_header_len:]
			}
			color.Danger.Printf("-------------------------------------------------\n")
			color.Danger.Printf(">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)\n", Ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			color.Danger.Printf("## bgs.protocol.rpc.Header\n")

			bgs_header := &protocol.Header{}
			if err := proto.Unmarshal(bgs_rpc_header_bytes, bgs_header); err != nil {
				color.Danger.Printf("Failed to parse address Header: %s\n", err)
			}
			PrintMessage(bgs_header)
			svc_hash := uint32(0)
			method_id := uint32(0)
			rcp_token := bgs_header.GetToken()
			rpc_kind := services.RPCCallKind(*bgs_header.ServiceId)
			color.Danger.Printf("* RPCKind: %s (service_id: %d)\n", services.RPCCallKind(*bgs_header.ServiceId), *bgs_header.ServiceId)
			color.Danger.Printf("+ Tracking Token: %d\n", rcp_token)
			if rpc_kind == "request" {
				svc_hash = bgs_header.GetServiceHash()
				method_id = bgs_header.GetMethodId()
				add_pending_response(TrackingToken(rcp_token), PendingResponse{svc_hash, method_id})
				color.Danger.Printf("  --> Added PendingResponse for Token %d\n", rcp_token)
			}
			if rpc_kind == "response" {
				pending, ok := recall_pending_response(TrackingToken(rcp_token))
				if ok {
					color.Danger.Printf("  <-- Recalled PendingResponse for Token %d\n", rcp_token)
					svc_hash = pending.ServiceHash
					method_id = pending.MethodId
				} else {
					color.Danger.Printf("  x-- Failed recalling a PendingReponse for Token %d\n", rcp_token)
					color.Danger.Printf("  x-- Will NOT decode message :(\n")
				}
			}
			if svc_hash > 0 {
				val, ok := services.Get(svc_hash)
				if ok {
					color.Danger.Printf("+ Service: %s (service_hash: %d)\n", val.Name(), svc_hash)
					if method_id > 0 {
						mval, mok := val.Method(uint16(method_id))
						if mok {
							color.Danger.Printf("+ Method: %s (method_id: %d)\n", mval, method_id)
							messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, *bgs_header.ServiceId))
							color.Danger.Printf("= MessageType : %s\n", messageName)
						} else {
							color.Danger.Printf("x Unknown method: %s\n", mval)
						}
					}
				} else {
					color.Danger.Printf("x Unknown Service hash: %04x\n", svc_hash)
				}
			}
			color.Danger.Printf("## bgs.protocol.rpc.Message\n")
			if bgs_rpc_message_len < 1 {
				color.Danger.Printf("[NO CONTENT]\n")
			} else {
				val, _ := services.Get(svc_hash)
				method, _ := val.Method(uint16(method_id))
				msg, err := services.ServiceMsg(val.Name(), method, *bgs_header.ServiceId, bgs_rpc_message_bytes)
				if err != nil {
					fmt.Println(err)
				} else {
					PrintMessage(*msg)
					mtype := protoreflect.FullName(services.PbMessageStr(val.Name(), method, *bgs_header.ServiceId))
					if fmt.Sprintf("%s", mtype) == "bgs.protocol.notification.v2.client.NotificationReceivedNotification" {
						ResolveNotificationPayload(bgs_rpc_message_bytes)
					}
				}
			}
			color.Danger.Printf("<<< EOF\n")
			Log_Mutex.Unlock()
			continue

		} else {
			data = nil
			data = make([]byte, 512)
			n, err := r.Read(data)
			_n := n
			_c := cap(data)
			for _n == _c {
				next := make([]byte, 512)
				nn, _ := r.Read(next)
				_n = nn
				n += _n
				data = append(data, next...)
				next = nil
			}
			if err != nil && err != io.EOF {
				color.Danger.Printf("unable to read data %v", err)
				break
			}
			if n > 0 {
				color.Danger.Printf("[conn_id = %d] from = %s (len = %d)\n", Ids, source, n)
				color.Danger.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					color.Danger.Printf("  -> Got Hanshake on %s side!\n", source)
				}
				c.Write(data)
				data = nil
				color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)
				continue
			}

		}

	}
}
