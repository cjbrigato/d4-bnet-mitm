package config

type Config struct {
	NoTLS               bool
	CertFile            string
	KeyFile             string
	ListenAddr          string
	RemoteAddr          string
	CheckRegistry       bool
	Verbose             bool
	CraftedAuthResponse bool
}

var (
	GlobalCfg = Config{}
)
