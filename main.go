package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"embed"

	"github.com/cjbrigato/d4-bnet-mitm/dynamic"
	"github.com/cjbrigato/d4-bnet-mitm/exchanges"
	"github.com/cjbrigato/d4-bnet-mitm/services"

	"github.com/pterm/pterm"
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
var ids int64 = 0
var Log_Mutex sync.RWMutex

var (
	noTLS         = flag.Bool("no-tls", false, "don't use TLS/SSL for both connections")
	certFile      = flag.String("cert", "", "X.509 certificate file")
	keyFile       = flag.String("keyfile", "", "X.509 key file")
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
	go exchanges.HandleServer(remote, client, "BGS::SERVER", exchanges.Ids)
	go exchanges.HandleClient(client, remote, "BGS::CLIENT", exchanges.Ids)

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

	//fmt.Printf("%s\n", banner)
	fbanner()
	fmt.Printf("                               %s%s", tlsMessage, "    +---Ready------------------ \n\n")
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
