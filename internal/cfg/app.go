package cfg

type ApplicationConfig struct {
	Dev           bool   `env:"DEV"`
	CdnURL        string `env:"CDN_URL"`
	SiteURL       string `env:"SITE_URL"`
	ImageProxyURL string `env:"IMGPROXY_URL"`
}
