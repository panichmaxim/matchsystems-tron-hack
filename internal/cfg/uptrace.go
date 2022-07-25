package cfg

type UptraceConfig struct {
	UptraceName string `env:"UPTRACE_NAME"`
	UptraceDSN  string `env:"UPTRACE_DSN"`
}
