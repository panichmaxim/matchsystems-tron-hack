package cfg

type MailgunConfig struct {
	MailgunDomain     string `env:"MAILGUN_DOMAIN"`
	MailgunPrivateKey string `env:"MAILGUN_PRIVATEKEY"`
	MailgunEndpoint   string `env:"MAILGUN_ENDPOINT"`
	MailgunFrom       string `env:"MAILGUN_FROM"`
}
