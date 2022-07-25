package cfg

type JwtConfig struct {
	//JwtSecret string `env:"JWT_SECRET,required"`
	JwtSecret string `env:"JWT_SECRET"`
}
