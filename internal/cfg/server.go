package cfg

type ServerConfig struct {
	ListenPort int `env:"LISTEN_PORT"`
}
