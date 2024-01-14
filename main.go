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
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"embed"

	"github.com/cjbrigato/d4-bnet-mitm/bnet/Fenris/ClientMessage"
	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol"
	"github.com/cjbrigato/d4-bnet-mitm/bnet/bgs/protocol/notification/v2/client"
	"github.com/cjbrigato/d4-bnet-mitm/certificate"
	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
	"github.com/gookit/color"
	"github.com/pterm/pterm"
	"golang.org/x/sys/windows"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protorange"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const banner = ` ____  _  _   _                _                  _ _             
|  _ \| || | | |__  _ __   ___| |_      _ __ ___ (_) |_ _ __ ___  
| | | | || |_| '_ \| '_ \ / _ \ __|____| '_ ` + "`" + ` _ \| | __| '_ ` + "`" + ` _ \ 
| |_| |__   _| |_) | | | |  __/ ||_____| | | | | | | |_| | | | | |
|____/   |_| |_.__/|_| |_|\___|\__|    |_| |_| |_|_|\__|_| |_| |_|
 [ TLS: using embedded X509 pair ]   +----Ready ----------------  `

//go:embed ssl/bnetserver.*
//go:embed ssl/Aurora*
//go:embed build/pb/*_bundle.binpb
var f embed.FS

var raddr string
var ids = 0
var log_mutex sync.RWMutex

var (
	noTLS         = flag.Bool("no-tls", false, "don't use TLS/SSL for both connections")
	certFile      = flag.String("cert", "", "X.509 certificate file")
	keyFile       = flag.String("keyfile", "", "X.509 key file")
	listenAddr    = flag.String("listen-addr", "127.0.0.1:1119", "Listen Address")
	remoteAddr    = flag.String("remote-addr", "185.60.112.74:1119", "Remote Address")
	checkRegistry = flag.Bool("check-registry", false, "check registry at startup for missing proto descriptor")
	verbose       = flag.Bool("verbose", false, "more verbose output with usefull messages in edge cases")
)

func fbanner() {
	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)
	str := banner
	strs := strings.Split(str, "")
	var fadeInfo string
	for i := 0; i < len(str); i++ {
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(fadeInfo)
}

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

/*func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] <listen-addr> <remote-addr>\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}*/

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
				log.Printf("unable to read frame: %v", err)
				break
			}

			LastReadFrame := []byte{}
			LastReadFrame = append(LastReadFrame, frame.Header.Bytes...)
			LastReadFrame = append(LastReadFrame, frame.Payload...)

			//_, err = c.Write(LastReadFrame)
			//color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)

			log_mutex.Lock()
			bgs_rpc_header_len := binary.BigEndian.Uint16(frame.Payload[0:])
			bgs_rpc_header_bytes := frame.Payload[2 : bgs_rpc_header_len+2]
			bgs_rpc_message_len := len(frame.Payload) - int(bgs_rpc_header_len+2)
			bgs_rpc_message_bytes := []byte{}
			if bgs_rpc_message_len > 0 {
				bgs_rpc_message_bytes = frame.Payload[2+bgs_rpc_header_len:]
			}

			log.Printf("-------------------------------------------------\n")
			log.Printf(">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)\n", ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			log.Printf("## bgs.protocol.rpc.Header\n")

			bgs_header := &protocol.Header{}
			if err := proto.Unmarshal(bgs_rpc_header_bytes, bgs_header); err != nil {
				log.Printf("Failed to parse address Header: %s\n", err)
			}
			PrintMessage(bgs_header)
			svc_hash := uint32(0)
			method_id := uint32(0)
			rcp_token := bgs_header.GetToken()
			rpc_kind := services.RPCCallKind(*bgs_header.ServiceId)
			log.Printf("* RPCKind: %s (service_id: %d)\n", services.RPCCallKind(*bgs_header.ServiceId), *bgs_header.ServiceId)
			log.Printf("+ Tracking Token: %d\n", rcp_token)
			if rpc_kind == "request" {
				svc_hash = bgs_header.GetServiceHash()
				method_id = bgs_header.GetMethodId()
				add_pending_response(TrackingToken(rcp_token), PendingResponse{svc_hash, method_id})
				log.Printf("  --> Added PendingResponse for Token %d\n", rcp_token)
			}
			if rpc_kind == "response" {
				pending, ok := recall_pending_response(TrackingToken(rcp_token))
				if ok {
					log.Printf("  <-- Recalled PendingResponse for Token %d\n", rcp_token)
					svc_hash = pending.ServiceHash
					method_id = pending.MethodId
				} else {
					log.Printf("  x-- Failed recalling a PendingReponse for Token %d\n", rcp_token)
					log.Printf("  x-- Will NOT decode message :(\n")
				}
			}
			if svc_hash > 0 {
				val, ok := services.Get(svc_hash)
				if ok {
					log.Printf("+ Service: %s (service_hash: %d)\n", val.Name(), svc_hash)
					if method_id > 0 {
						mval, mok := val.Method(uint16(method_id))
						if mok {
							log.Printf("+ Method: %s (method_id: %d)\n", mval, method_id)
							messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, *bgs_header.ServiceId))
							log.Printf("= MessageType : %s\n", messageName)
						} else {
							log.Printf("x Unknown method: %s\n", mval)
						}
					}
				} else {
					log.Printf("x Unknown Service hash: %04x\n", svc_hash)
				}
			}
			log.Printf("## bgs.protocol.rpc.Message\n")
			if bgs_rpc_message_len < 1 {
				log.Printf("[NO CONTENT]\n")
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
						ResolveNotificationPayload(&bgs_rpc_message_bytes)
					}
				}
			}

			log.Printf("<<< EOF\n")
			log_mutex.Unlock()

			bgs_rpc_message_len = len(bgs_rpc_message_bytes)
			header_len := uint16(len(bgs_rpc_header_bytes))
			header_len_bytes := make([]byte, 2)
			binary.BigEndian.PutUint16(header_len_bytes, header_len)

			newframepayload := make([]byte, len(header_len_bytes)+len(bgs_rpc_header_bytes)+len(bgs_rpc_message_bytes))
			newframepayload = append(header_len_bytes, bgs_rpc_header_bytes...)
			if bgs_rpc_message_len > 0 {
				newframepayload = append(newframepayload, bgs_rpc_message_bytes...)
			}

			newframe := append(frame.Header.Bytes, newframepayload...)
			//newframe := ws.NewBinaryFrame(newframepayload)
			//ws.WriteFrame(c, newframe)
			/*fmt.Printf("------------\n")
			fmt.Printf("LastReadFrame: \n,%s", hex.Dump(LastReadFrame))
			fmt.Printf("=================\n")
			fmt.Printf("newframe: \n,%s", hex.Dump(newframe))*/
			c.Write(newframe)
			//newframe :=
			//_, err = c.Write(LastReadFrame)
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
				log.Printf("unable to read data %v", err)
				break
			}
			if n > 0 {
				log.Printf("[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
				log.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					log.Printf("  -> Got Hanshake on %s side!\n", source)
				}
				c.Write(data)
				data = nil
				//color.Infoln("Wrote %d bytes representing last frame from remote conn to cli conn", n)
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
				log.Printf("unable to read frame: %v", err)
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

			log.Printf("-------------------------------------------------\n")
			log.Printf(">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)\n", ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			log.Printf("## bgs.protocol.rpc.Header\n")

			bgs_header := &protocol.Header{}
			if err := proto.Unmarshal(bgs_rpc_header_bytes, bgs_header); err != nil {
				log.Printf("Failed to parse address Header: %s\n", err)
			}
			PrintMessage(bgs_header)
			svc_hash := uint32(0)
			method_id := uint32(0)
			rcp_token := bgs_header.GetToken()
			rpc_kind := services.RPCCallKind(*bgs_header.ServiceId)
			log.Printf("* RPCKind: %s (service_id: %d)\n", services.RPCCallKind(*bgs_header.ServiceId), *bgs_header.ServiceId)
			log.Printf("+ Tracking Token: %d\n", rcp_token)
			if rpc_kind == "request" {
				svc_hash = bgs_header.GetServiceHash()
				method_id = bgs_header.GetMethodId()
				add_pending_response(TrackingToken(rcp_token), PendingResponse{svc_hash, method_id})
				log.Printf("  --> Added PendingResponse for Token %d\n", rcp_token)
			}
			if rpc_kind == "response" {
				pending, ok := recall_pending_response(TrackingToken(rcp_token))
				if ok {
					log.Printf("  <-- Recalled PendingResponse for Token %d\n", rcp_token)
					svc_hash = pending.ServiceHash
					method_id = pending.MethodId
				} else {
					log.Printf("  x-- Failed recalling a PendingReponse for Token %d\n", rcp_token)
					log.Printf("  x-- Will NOT decode message :(\n")
				}
			}
			if svc_hash > 0 {
				val, ok := services.Get(svc_hash)
				if ok {
					log.Printf("+ Service: %s (service_hash: %d)\n", val.Name(), svc_hash)
					if method_id > 0 {
						mval, mok := val.Method(uint16(method_id))
						if mok {
							log.Printf("+ Method: %s (method_id: %d)\n", mval, method_id)
							messageName := protoreflect.FullName(services.PbMessageStr(val.Name(), mval, *bgs_header.ServiceId))
							log.Printf("= MessageType : %s\n", messageName)
						} else {
							log.Printf("x Unknown method: %s\n", mval)
						}
					}
				} else {
					log.Printf("x Unknown Service hash: %04x\n", svc_hash)
				}
			}
			log.Printf("## bgs.protocol.rpc.Message\n")
			if bgs_rpc_message_len < 1 {
				color.Info.Printf("[NO CONTENT]\n")
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
						ResolveNotificationPayload(&bgs_rpc_message_bytes)
					}
				}
			}
			color.Info.Printf("<<< EOF\n")
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
				log.Printf("unable to read data %v", err)
				break
			}
			if n > 0 {
				log.Printf("[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
				log.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					log.Printf("  -> Got Hanshake on %s side!\n", source)
				}
				data = nil
				continue
			}

		}

	}
}

func ResolveNotificationPayload(bgs_rpc_message_bytes *[]byte) {
	notifmsg := &client.NotificationReceivedNotification{}
	if err := proto.Unmarshal(*bgs_rpc_message_bytes, notifmsg); err != nil {
		fmt.Printf("Failed to parse Notification: %s\n", err)
	}
	notif := notifmsg.GetNotifications()[0].GetAttribute()
	messageid := notif[0].GetValue().GetIntValue()
	payload := notif[1].GetValue().GetBlobValue()
	var messageid_type string
	switch messageid {
	case 0:
		messageid_type = "Fenris.ClientMessage.FindUserProxyResponse"
		findUserProxyMsg := &ClientMessage.FindUserProxyResponse{}
		if err := proto.Unmarshal(payload, findUserProxyMsg); err != nil {
			log.Printf("Failed to parse FindUserProxyResponse: %s\n", err)
		}
		var err error
		findUserProxyMsg.ConnectInfo.Address = proto.String("127.0.0.1:61000")
		findUserProxyMsg.ConnectInfo.Port = proto.Uint32(61000)
		notifmsg.GetNotifications()[0].GetAttribute()[1].Value.BlobValue, err = proto.Marshal(findUserProxyMsg)
		if err != nil {
			log.Printf("Failed to Marshal crafted FindUserProxyResponse: %s\n", err)
		}
		*bgs_rpc_message_bytes, err = proto.Marshal(notifmsg)
		if err != nil {
			log.Printf("Failed to Marshal crafted : %s\n", err)
		}
		log.Printf("### %s (as FEN.NotificationMessage.Payload)\n", messageid_type)
		PrintMessage(findUserProxyMsg)
	case 5:
		messageid_type = "Fenris.ClientMessage.PingConnectInfoSingleResult"
		FenrisMessage := &ClientMessage.PingConnectInfoSingleResult{}
		if err := proto.Unmarshal(payload, FenrisMessage); err != nil {
			log.Printf("Failed to parse FindUserProxyResponse: %s\n", err)
		}
		log.Printf("### %s (as FEN.NotificationMessage.Payload)\n", messageid_type)
		PrintMessage(FenrisMessage)
	case 2:
		messageid_type = "Fenris.ClientMessage.QueueUpdate"
		FenrisMessage := &ClientMessage.QueueUpdate{}
		if err := proto.Unmarshal(payload, FenrisMessage); err != nil {
			log.Printf("Failed to parse FindUserProxyResponse: %s\n", err)
		}
		log.Printf("### %s (as FEN.NotificationMessage.Payload)\n", messageid_type)
		PrintMessage(FenrisMessage)
	default:
		messageid_type = "unknown"
		FenrisMessage, err := dynamic.ParseAs(messageid_type, payload)
		if err != nil {
			fmt.Printf("Failed to parse %s: %s\n", messageid_type, err)
			log.Printf("### %s (as FEN.NotificationMessage.Payload)\n", messageid_type)
			PrintMessage(*FenrisMessage)
		} else {
			fmt.Printf("%s\n", hex.Dump(payload))
		}
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
	go handleServer(remote, client, "BGS::SERVER", ids)

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

/*
	func amAdmin() bool {
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		if err != nil {
			fmt.Println("admin no")
			return false
		}
		fmt.Println("admin yes")
		return true
	}
*/
func amAdmin() bool {
	elevated := windows.GetCurrentProcessToken().IsElevated()
	//fmt.Printf("admin %v\n", elevated)
	return elevated
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func permitExclusion() {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	script := fmt.Sprintf(`Add-MpPreference -ExclusionProcess "%s"`, filepath.Base(executable))
	script2 := fmt.Sprintf(`Add-MpPreference -ControlledFolderAccessAllowedApplications "%s"`, executable)

	_, err = exec.Command("powershell.exe", script).Output()
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("powershell.exe", script2).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success")

}

func main() {

	if !amAdmin() {
		runMeElevated()
		time.Sleep(10 * time.Second)
	}

	flag.Parse()

	permitExclusion()
	init_pending_responses()
	dynamic.Register("build/pb/bgs_bundle.binpb", &f)
	dynamic.Register("build/pb/fenris_bundle.binpb", &f)
	if *checkRegistry {
		services.Test_protos()
	}
	certificate.InstallCertificate("ssl/AuroraCA.cer", &f)
	certificate.InstallCertificate("ssl/AuroraRootCA.cer", &f)
	certificate.InstallCertificate("ssl/bnetserver.crt", &f)

	var err error

	laddr, err := net.ResolveTCPAddr("tcp", *listenAddr)
	if err != nil {
		fatal("listen addr: %s\n", err)
	}

	raddr = *remoteAddr

	var ln net.Listener

	ln, err = net.ListenTCP("tcp", laddr)
	if err != nil {
		fatal("listen:", err, "\n")
	}

	if !*noTLS {
		config := tls.Config{}
		config.Certificates = make([]tls.Certificate, 1)
		if !*noTLS && !(exists(*certFile) && exists(*keyFile)) {
			config.Certificates[0], err = LoadEmbededX509KeyPair()
		} else {
			config.Certificates[0], err = tls.LoadX509KeyPair(*certFile, *keyFile)
		}

		if err != nil {
			fatal("loading cert:", err)
		}

		ln = tls.NewListener(ln, &config)
	}

	fbanner()
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
