package main

import (
	"flag"

	"embed"

	"github.com/cjbrigato/d4-bnet-mitm/bgspacket"
	"github.com/cjbrigato/d4-bnet-mitm/config"
	"github.com/cjbrigato/d4-bnet-mitm/network"
	"github.com/cjbrigato/d4-bnet-mitm/ui"
)

//go:embed network/ssl/bnetserver.*
//go:embed network/ssl/Aurora*
//go:embed build/pb/*_bundle.binpb
var f embed.FS

var (
	noTLS               = flag.Bool("no-tls", false, "don't use TLS/SSL for both connections")
	certFile            = flag.String("cert", "", "X.509 certificate file")
	keyFile             = flag.String("keyfile", "", "X.509 key file")
	listenAddr          = flag.String("listen-addr", "127.0.0.1:1119", "Listen Address")
	remoteAddr          = flag.String("remote-addr", "185.60.112.74:1119", "Remote Address")
	checkRegistry       = flag.Bool("check-registry", false, "check registry at startup for missing proto descriptor")
	verbose             = flag.Bool("verbose", false, "more verbose output with usefull messages in edge cases")
	craftedAuthResponse = flag.Bool("crafted-auth-response", false, "send a crafted auth response to the client")
)

func main() {

	flag.Parse()
	config.GlobalCfg = config.Config{
		NoTLS:               *noTLS,
		CertFile:            *certFile,
		KeyFile:             *keyFile,
		ListenAddr:          *listenAddr,
		RemoteAddr:          *remoteAddr,
		CheckRegistry:       *checkRegistry,
		Verbose:             *verbose,
		CraftedAuthResponse: *craftedAuthResponse,
	}

	ui.Preamble()

	bgspacket.ProtocolInit(*checkRegistry, &f)
	network.NetworkInit(&f)

	ui.PreStart()
	ui.StartUi()
}
