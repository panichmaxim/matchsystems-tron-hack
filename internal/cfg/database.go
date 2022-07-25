package cfg

type DatabaseConfig struct {
	VerboseSQL     bool   `env:"VERBOSE_SQL"`
	DatabaseURL    string `env:"DATABASE_URL"`
	LoadFixtures   bool   `env:"LOAD_FIXTURES"`
	LoadMigrations bool   `env:"LOAD_MIGRATIONS"`
}
