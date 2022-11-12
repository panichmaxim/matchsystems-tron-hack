package cfg

type DistributedConfig struct {
	App        ApplicationConfig
	Kubernetes KubernetesPodInfoConfig
	Database   DatabaseConfig
	BtcNeo     BtcNeo4jConfig
	EthNeo     EthNeo4jConfig
	TronNeo    TronNeo4jConfig
	Elastic    ElasticConfig
	Uptrace    UptraceConfig
	Mailgun    MailgunConfig
	Smtp       SmtpConfig
	Sentry     SentryConfig
	Server     ServerConfig
	Jwt        JwtConfig
	Telegram   TelegramConfig
}
