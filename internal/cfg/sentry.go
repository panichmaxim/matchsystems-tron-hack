package cfg

type SentryConfig struct {
	SentryDSN string `env:"SENTRY_DSN"`
}
