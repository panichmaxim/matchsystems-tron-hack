package cfg

type TelegramConfig struct {
	ChatID   string `env:"TELEGRAM_CHAT_ID"`
	BotToken string `env:"TELEGRAM_BOT_TOKEN"`
}
