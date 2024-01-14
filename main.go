package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"embed"

	"github.com/cjbrigato/d4-bnet-mitm/bgspacket"
	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	log2 "github.com/cjbrigato/d4-bnet-mitm/log"
	"github.com/cjbrigato/d4-bnet-mitm/services"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
	"github.com/pterm/pterm"
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

func isServerSource(source string) bool {
	return strings.Contains(source, "BGS::SERVER")
}

func handleFrame(r io.Reader, c io.Writer, source string, conn_id int) {

	upgraded := false
	data := make([]byte, 512)
	for {
		if upgraded {
			frame, err := ws.ReadFrame(r)
			if err != nil {
				log2.Error(nil, "unable to read frame: %v", err)
				break
			}

			log_mutex.Lock()
			log2.Info(nil, "-------------------------------------------------")
			log2.Info(nil, ">>> WS::FRAME -> %d:%s (%s, fin = %t, %d bytes)", ids, source, frame.Header.OpCode, frame.Header.Fin, frame.Header.Length)
			bgs_packet := bgspacket.NewBgsPacketFromFrame(&frame, false)
			log2.Info(nil, "<<< EOF")
			log_mutex.Unlock()

			if isServerSource(source) {
				c.Write(bgs_packet.WSEncapsulate(frame))
			}
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
				log2.Error(nil, "unable to read data %v", err)
				break
			}
			if n > 0 {
				log.Printf("[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
				log.Printf("  -> Pre-upgrade, Waiting for Handshake\n")
				if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
					upgraded = true
					log2.Info(nil, "  -> Got Hanshake on %s side!\n", source)
				}
				if isServerSource(source) {
					c.Write(data)
				}
				data = nil
				continue
			}

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
		log2.Error(nil, "dial remote: %v", err)
		return
	}

	defer remote.Close()
	go handleFrame(remote, client, "BGS::SERVER", ids)

	rs, ws := io.Pipe()
	ts := io.MultiWriter(remote, ws)
	go handleFrame(rs, remote, "BGS::CLIENT", ids)
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
			log2.Info(nil, "new connection")
			logproxy(c)
			done <- c
		}
	}()
	return ch
}

func main() {

	flag.Parse()

	InstallCertificates()
	dynamic.Register("build/pb/bgs_bundle.binpb", &f)
	dynamic.Register("build/pb/fenris_bundle.binpb", &f)
	if *checkRegistry {
		services.Test_protos()
	}

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
	log2.Info(nil, "::  Listening on <- %s", laddr)
	log2.Info(nil, ":: with upstream -> %s", raddr)

	done := cleanup()
	conns := handle(done)

	for {
		c, err := ln.Accept()
		if err != nil {
			log2.Debug(nil, "accept: %q", err)
			continue
		}
		conns <- c
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
