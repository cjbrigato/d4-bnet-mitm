package network

import (
	"bytes"
	"crypto/tls"
	"embed"
	"fmt"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cjbrigato/d4-bnet-mitm/bgspacket"
	"github.com/cjbrigato/d4-bnet-mitm/config"
	log2 "github.com/cjbrigato/d4-bnet-mitm/log"
	"github.com/cjbrigato/d4-bnet-mitm/ui"
	"github.com/cjbrigato/d4-bnet-mitm/ws"
)

var raddr string
var ids = 0
var jids = make(map[int]int)
var had_first_byte bool
var started_at time.Time

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func isServerSource(source string) bool {
	return strings.Contains(source, "BGS::SERVER")
}

func SourceFromString(s string) bgspacket.PacketSource {
	switch s {
	case "BGS::CLIENT":
		return bgspacket.SRC_CLIENT
	case "BGS::SERVER":
		return bgspacket.SRC_SERVER
	}
	return bgspacket.SRC_UNKNOWN
}

func handleFrame(r io.Reader, c io.Writer, source string, conn_id int) {
	upgraded := false

	data := make([]byte, 512)
	for !upgraded {
		if err := EnsureUpgrade(data, r, source, &upgraded, c); err != nil {
			log2.Error(nil, "unable to upgrade: %v", err)
			panic(err)
		}
	}
	for {
		frame, err := ws.ReadFrame(r)
		if err != nil {
			log2.Error(nil, "unable to read frame: %v", err)
			break
		}

		log2.Wmutex.Lock()
		bgs_packet := bgspacket.NewBgsPacketFromFrame(&frame, SourceFromString(source), config.GlobalCfg.CraftedAuthResponse)
		if !had_first_byte {
			started_at = time.Now()
			had_first_byte = true
		}
		jids[ids]++
		ui.App.QueueUpdateDraw(func() {
			src_color := "[green]CLIENT[blue] -->[white]"
			if isServerSource(source) {
				src_color = "[orange]<-- [darkcyan]SERVER[white]"
			}
			var btype string
			var lastbtype string
			var PaddedMessageTypeString string
			var isError bool
			token := bgs_packet.RpcToken
			sname, _ := strings.CutPrefix(fmt.Sprintf("%s", bgs_packet.ServiceName), "bgs.protocol.")
			if bgs_packet.ErrorString == "" {
				btype, _ = strings.CutPrefix(fmt.Sprintf("%s", bgs_packet.MessageType), "bgs.protocol.")
				lastbtype = btype[strings.LastIndex(btype, ".")+1:]
			} else {
				btype = "Error: see packet details"
				lastbtype = btype
				isError = true
			}
			PaddedMessageTypeString = ui.PaddedMessageTypeString(lastbtype)
			kind := bgs_packet.RpcKind
			if kind == "request" {
				kind = "request "
			}

			if !isError {
				PaddedMessageTypeString = fmt.Sprintf("[green]%s[white]", PaddedMessageTypeString)
			} else {
				PaddedMessageTypeString = fmt.Sprintf("[red]%s[white]", PaddedMessageTypeString)
			}

			ui.PacketList.AddItem(fmt.Sprintf("[grey]%4d.[white] [blue]%8.3f[white]  %s [white] %3d %s  %s  %s", jids[ids], float32(time.Since(started_at).Milliseconds())/1000.0, src_color, token, kind, PaddedMessageTypeString, sname), "", 0,
				func() {
					ui.App.SetFocus(ui.PacketInfo)
				})
		})
		/*switch bgs_packet.MessageType {
		case "bgs.protocol.game_utilities.v2.client.ProcessTaskResponse":
			bgs_packet.BGSHeader.PrintMessage()
			if len(bgs_packet.BGSMessage.Bytes) > 0 {
				bgs_packet.BGSMessage.PrintMessage()
			}
		case "bgs.protocol.game_utilities.v2.client.ProcessTaskRequest":
			bgs_packet.BGSHeader.PrintMessage()
			if len(bgs_packet.BGSMessage.Bytes) > 0 {
				bgs_packet.BGSMessage.PrintMessage()
			}
		}
		log2.Info(nil, "<<< EOF")*/
		log2.Wmutex.Unlock()

		if isServerSource(source) {
			c.Write(bgs_packet.WSEncapsulate(frame))
		}
	}
}

func EnsureUpgrade(data []byte, r io.Reader, source string, upgraded *bool, c io.Writer) error {
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
		return err
	}
	if n > 0 {
		log2.Info(nil, "[conn_id = %d] from = %s (len = %d)\n", ids, source, n)
		log2.Info(nil, "  -> Pre-upgrade, Waiting for Handshake\n")
		if bytes.Contains(data, []byte("v1.rpc.battle.net")) {
			*upgraded = true
			log2.Info(nil, "  -> Got Hanshake on %s side!\n", source)
		}
		if isServerSource(source) {
			c.Write(data)
		}
		data = nil
	}
	return nil
}

func logproxy(client net.Conn) {

	var (
		remote net.Conn
		err    error
	)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	if !config.GlobalCfg.NoTLS {
		remote, err = tls.Dial("tcp", config.GlobalCfg.RemoteAddr, conf)
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
			log2.Info(nil, "closing connection")
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

func LoadEmbededX509KeyPair(f *embed.FS) (tls.Certificate, error) {
	certPEMBlock, err := f.ReadFile("vfs/bnetserver.crt")
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEMBlock, err := f.ReadFile("vfs/bnetserver.key")
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
}

func NetworkInit(f *embed.FS) {

	if runtime.GOOS == "windows" {
		InstallCertificates(f)
	}

	var err error

	laddr, err := net.ResolveTCPAddr("tcp", config.GlobalCfg.ListenAddr)
	if err != nil {
		log2.Fatal(nil, "listen addr: %s", err)
	}

	raddr = config.GlobalCfg.RemoteAddr

	var ln net.Listener

	ln, err = net.ListenTCP("tcp", laddr)
	if err != nil {
		log2.Fatal(nil, "listen: %v", err)
	}

	if !config.GlobalCfg.NoTLS {
		tlsconfig := tls.Config{}
		tlsconfig.Certificates = make([]tls.Certificate, 1)
		if !config.GlobalCfg.NoTLS && !(exists(config.GlobalCfg.CertFile) && exists(config.GlobalCfg.KeyFile)) {
			tlsconfig.Certificates[0], err = LoadEmbededX509KeyPair(f)
		} else {
			tlsconfig.Certificates[0], err = tls.LoadX509KeyPair(config.GlobalCfg.CertFile, config.GlobalCfg.KeyFile)
		}

		if err != nil {
			log2.Fatal(nil, "loading cert: %v", err)
		}

		ln = tls.NewListener(ln, &tlsconfig)
	}

	done := cleanup()
	conns := handle(done)

	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				log2.Debug(nil, "accept: %q", err)
				continue
			}
			conns <- c
		}
	}()

	log2.Info(nil, "::  Listening on <- %s", laddr)
	log2.Info(nil, ":: with upstream -> %s", raddr)
}
