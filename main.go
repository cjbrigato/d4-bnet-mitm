package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/notification/v2/client"
	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
	"github.com/gookit/color"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"

	"embed"
)

const banner = ` ____  _  _   _                _                  _ _             
|  _ \| || | | |__  _ __   ___| |_      _ __ ___ (_) |_ _ __ ___  
| | | | || |_| '_ \| '_ \ / _ \ __|____| '_ ` + "`" + ` _ \| | __| '_ ` + "`" + ` _ \ 
| |_| |__   _| |_) | | | |  __/ ||_____| | | | | | | |_| | | | | |
|____/   |_| |_.__/|_| |_|\___|\__|    |_| |_| |_|_|\__|_| |_| |_|`

//go:embed ssl/bnetserver.*
//go:embed build/pb/*_bundle.binpb
var f embed.FS

var raddr string
var ids = 0
var log_mutex sync.RWMutex

var (
	noTLS         = flag.Bool("no-tls", false, "don't use TLS/SSL for both connections")
	certFile      = flag.String("cert", "", "X.509 certificate file")
	keyFile       = flag.String("keyfile", "", "X.509 key file")
	checkRegistry = flag.Bool("check-registry", false, "check registry at startup for missing proto descriptor")
	verbose       = flag.Bool("verbose", false, "more verbose output with usefull messages in edge cases")
)

func LoadEmbededX509KeyPair() (tls.Certificate, error) {
	certPEMBlock, err := f.ReadFile("ssl/bnetserver.crt")
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEMBlock, err := f.ReadFile("ssl/bnetserver.key")
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] <listen-addr> <remote-addr>\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
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

var pending_mutex sync.RWMutex

type PendingResponse struct {
	ServiceHash uint32
	MethodId    uint32
}
type TrackingToken uint32

var pending_responses map[TrackingToken]PendingResponse

func init_pending_responses() {
	pending_responses = make(map[TrackingToken]PendingResponse)
}
func add_pending_response(token TrackingToken, pending_spec PendingResponse) {
	pending_mutex.Lock()
	pending_responses[TrackingToken(token)] = pending_spec
	pending_mutex.Unlock()
}
func recall_pending_response(token TrackingToken) (PendingResponse, bool) {
	pending_mutex.Lock()
	pending, ok := pending_responses[token]
	if ok {
		delete(pending_responses, token)
	}
	pending_mutex.Unlock()
	return pending, ok
}

func handleServer(r io.Reader, c io.Writer, source string, conn_id int) {

	upgraded := false
	data := make([]byte, 512)
	for {

		if upgraded {
			frame, err := ws.ReadFrame(r)
			if err != nil {
				fmt.Printf("unable to read frame: %v", err)
				break
			}

			LastReadFrame := []byte{}
			LastReadFrame = append(LastReadFrame, frame.Header.Bytes...)
			LastReadFrame = append(LastReadFrame, frame.Payload...)

			n, err := c.Write(LastReadFrame)
			color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)

			log_mutex.Lock()
			bgs_rpc_header_len := binary.BigEndian.Uint16(frame.Payload[0:])
			bgs_rpc_header_bytes := frame.Payload[2 : bgs_rpc_header_len+2]
			bgs_rpc_message_len := len(frame.Payload) - int(bgs_rpc_header_len+2)
			bgs_rpc_message_bytes := []byte{'0'}
			if bgs_rpc_message_len > 0 {
				bgs_rpc_message_bytes = frame.Payload[2+bgs_rpc_header_len:]
			}

			fmt.Printf("-------------------------------------------------\n")
			fmt.Printf(">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)\n", ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			fmt.Printf("## bgs.protocol.rpc.Header\n")

			bgs_header := &protocol.Header{}
			if err := proto.Unmarshal(bgs_rpc_header_bytes, bgs_header); err != nil {
				fmt.Printf("Failed to parse address Header: %s\n", err)
			}
			PrintMessage(bgs_header)
			svc_hash := uint32(0)
			method_id := uint32(0)
			rcp_token := bgs_header.GetToken()
			rpc_kind := services.RPCCallKind(*bgs_header.ServiceId)
			fmt.Printf("* RPCKind: %s (service_id: %d)\n", services.RPCCallKind(*bgs_header.ServiceId), *bgs_header.ServiceId)
			fmt.Printf("+ Tracking Token: %d\n", rcp_token)
			if rpc_kind == "request" {
				svc_hash = bgs_header.GetServiceHash()
				method_id = bgs_header.GetMethodId()
				add_pending_response(TrackingToken(rcp_token), PendingResponse{svc_hash, method_id})
				fmt.Printf("  --> Added PendingResponse for Token %d\n", rcp_token)
			}
			if rpc_kind == "response" {
				pending, ok := recall_pending_response(TrackingToken(rcp_token))
				if ok {
					fmt.Printf("  <-- Recalled PendingResponse for Token %d\n", rcp_token)
					svc_hash = pending.ServiceHash
					method_id = pending.MethodId
				} else {
					fmt.Printf("  x-- Failed recalling a PendingReponse for Token %d\n", rcp_token)
					fmt.Printf("  x-- Will NOT decode message :(\n")
				}
			}
			if svc_hash > 0 {
				val, ok := services.Get(svc_hash)
				if ok {
					fmt.Printf("+ Service: %s (service_hash: %d)\n", val.Name(), svc_hash)
					if method_id > 0 {
						mval, mok := val.Method(uint16(method_id))
						if mok {
							fmt.Printf("+ Method: %s (method_id: %d)\n", mval, method_id)
							messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, *bgs_header.ServiceId))
							fmt.Printf("= MessageType : %s\n", messageName)
						} else {
							fmt.Printf("x Unknown method: %s\n", mval)
						}
					}
				} else {
					fmt.Printf("x Unknown Service hash: %04x\n", svc_hash)
				}
			}
			fmt.Printf("## bgs.protocol.rpc.Message\n")
			if bgs_rpc_message_len < 1 {
				fmt.Printf("[NO CONTENT]\n")
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
			fmt.Printf("<<< EOF\n")
			log_mutex.Unlock()
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
				fmt.Printf("unable to read data %v", err)
				break
			}
			if n > 0 {
				fmt.Printf("[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
				fmt.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					fmt.Printf("  -> Got Hanshake on %s side!\n", source)
				}
				c.Write(data)
				data = nil
				color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)
				continue
			}
		}
	}
}

func dumpData(r io.Reader, source string, conn_id int) {

	upgraded := false
	data := make([]byte, 512)
	for {

		if upgraded {
			frame, err := ws.ReadFrame(r)
			if err != nil {
				fmt.Printf("unable to read frame: %v", err)
				break
			}
			log_mutex.Lock()
			bgs_rpc_header_len := binary.BigEndian.Uint16(frame.Payload[0:])
			bgs_rpc_header_bytes := frame.Payload[2 : bgs_rpc_header_len+2]
			bgs_rpc_message_len := len(frame.Payload) - int(bgs_rpc_header_len+2)
			bgs_rpc_message_bytes := []byte{'0'}
			if bgs_rpc_message_len > 0 {
				bgs_rpc_message_bytes = frame.Payload[2+bgs_rpc_header_len:]
			}

			fmt.Printf("-------------------------------------------------\n")
			fmt.Printf(">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)\n", ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			fmt.Printf("## bgs.protocol.rpc.Header\n")

			bgs_header := &protocol.Header{}
			if err := proto.Unmarshal(bgs_rpc_header_bytes, bgs_header); err != nil {
				fmt.Printf("Failed to parse address Header: %s\n", err)
			}
			PrintMessage(bgs_header)
			svc_hash := uint32(0)
			method_id := uint32(0)
			rcp_token := bgs_header.GetToken()
			rpc_kind := services.RPCCallKind(*bgs_header.ServiceId)
			fmt.Printf("* RPCKind: %s (service_id: %d)\n", services.RPCCallKind(*bgs_header.ServiceId), *bgs_header.ServiceId)
			fmt.Printf("+ Tracking Token: %d\n", rcp_token)
			if rpc_kind == "request" {
				svc_hash = bgs_header.GetServiceHash()
				method_id = bgs_header.GetMethodId()
				add_pending_response(TrackingToken(rcp_token), PendingResponse{svc_hash, method_id})
				fmt.Printf("  --> Added PendingResponse for Token %d\n", rcp_token)
			}
			if rpc_kind == "response" {
				pending, ok := recall_pending_response(TrackingToken(rcp_token))
				if ok {
					fmt.Printf("  <-- Recalled PendingResponse for Token %d\n", rcp_token)
					svc_hash = pending.ServiceHash
					method_id = pending.MethodId
				} else {
					fmt.Printf("  x-- Failed recalling a PendingReponse for Token %d\n", rcp_token)
					fmt.Printf("  x-- Will NOT decode message :(\n")
				}
			}
			if svc_hash > 0 {
				val, ok := services.Get(svc_hash)
				if ok {
					fmt.Printf("+ Service: %s (service_hash: %d)\n", val.Name(), svc_hash)
					if method_id > 0 {
						mval, mok := val.Method(uint16(method_id))
						if mok {
							fmt.Printf("+ Method: %s (method_id: %d)\n", mval, method_id)
							messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, *bgs_header.ServiceId))
							fmt.Printf("= MessageType : %s\n", messageName)
						} else {
							fmt.Printf("x Unknown method: %s\n", mval)
						}
					}
				} else {
					fmt.Printf("x Unknown Service hash: %04x\n", svc_hash)
				}
			}
			fmt.Printf("## bgs.protocol.rpc.Message\n")
			if bgs_rpc_message_len < 1 {
				fmt.Printf("[NO CONTENT]\n")
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
			fmt.Printf("<<< EOF\n")
			log_mutex.Unlock()
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
				fmt.Printf("unable to read data %v", err)
				break
			}
			if n > 0 {
				fmt.Printf("[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
				fmt.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					fmt.Printf("  -> Got Hanshake on %s side!\n", source)
				}
				data = nil
				continue
			}
		}
	}
}

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

func logproxy(client net.Conn) {

	var (
		remote net.Conn
		err    error
	)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	if !*noTLS {
		remote, err = tls.Dial("tcp", raddr, conf)
	} else {
		remote, err = net.Dial("tcp", raddr)
	}
	if err != nil {
		log.Println("dial remote:", err)
		return
	}

	defer remote.Close()

	//rc, wc := io.Pipe()
	//tc := io.MultiWriter(client, wc)
	go handleServer(remote, client, "BGS::SERVER", ids)
	//go func() {
	//	io.Copy(tc, remote)
	//}()

	rs, ws := io.Pipe()
	ts := io.MultiWriter(remote, ws)
	go dumpData(rs, "BGS::CLIENT", ids)
	io.Copy(ts, client)
}

func cleanup() chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for conn := range ch {
			log.Println("closing connection")
			conn.Close()
		}
	}()
	return ch
}

func handle(done chan<- net.Conn) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for c := range ch {
			log.Println("new connection")
			logproxy(c)
			done <- c
		}
	}()
	return ch
}

func main() {

	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}

	init_pending_responses()
	dynamic.Register("build/pb/bgs_bundle.binpb", &f)
	dynamic.Register("build/pb/fenris_bundle.binpb", &f)
	if *checkRegistry {
		services.Test_protos()
	}

	var err error

	laddr, err := net.ResolveTCPAddr("tcp", flag.Arg(0))
	if err != nil {
		fatal("listen addr: %s\n", err)
	}

	raddr = flag.Arg(1)

	var ln net.Listener

	ln, err = net.ListenTCP("tcp", laddr)
	if err != nil {
		fatal("listen:", err, "\n")
	}

	tlsMessage := " [ TLS: using embedded X509 pair ]"
	if !*noTLS {
		config := tls.Config{}
		config.Certificates = make([]tls.Certificate, 1)
		if !*noTLS && !(exists(*certFile) && exists(*keyFile)) {
			config.Certificates[0], err = LoadEmbededX509KeyPair()
		} else {
			tlsMessage = "[ TLS: Using CLI specified key/cert files ]"
			config.Certificates[0], err = tls.LoadX509KeyPair(*certFile, *keyFile)
		}

		if err != nil {
			fatal("loading cert:", err)
		}

		ln = tls.NewListener(ln, &config)
	}

	fmt.Printf("%s\n", banner)
	fmt.Printf(" %s%s", tlsMessage, "    +---Ready------------------ \n\n")
	log.Println("::  Listening on <-", laddr)
	log.Println(":: with upstream ->", raddr)

	done := cleanup()
	conns := handle(done)

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		conns <- c
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
