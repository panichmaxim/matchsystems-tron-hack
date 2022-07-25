package cfg

type BtcNeo4jConfig struct {
	Username string `env:"BTC_NEO4J_USERNAME"`
	Password string `env:"BTC_NEO4J_PASSWORD"`
	Address  string `env:"BTC_NEO4J_ADDRESS"`
}

type EthNeo4jConfig struct {
	Username string `env:"ETH_NEO4J_USERNAME"`
	Password string `env:"ETH_NEO4J_PASSWORD"`
	Address  string `env:"ETH_NEO4J_ADDRESS"`
}
