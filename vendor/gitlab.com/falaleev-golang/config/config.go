package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/hashicorp/go-multierror"
	"github.com/joho/godotenv"
	"os"
)

func loadStruct(cfg interface{}, fileNames ...string) error {
	var valid []string
	for _, f := range fileNames {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			continue
		}
		valid = append(valid, f)
	}

	_ = godotenv.Overload(valid...)

	return env.Parse(cfg)
}

// Load обеспечивает загрузку .env файлов в переданном порядке в первом аргументе и перезаписывает
// значения из последующих файлов
func Load(files []string, cfg ...interface{}) error {
	var result error

	for _, i := range cfg {
		if err := loadStruct(i, files...); err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result
}

// LoadDefault обеспечивает загрузку .env, .env.local файлов в переданном порядке в первом аргументе
// и перезаписывает значения из последующих файлов
func LoadDefault(cfg ...interface{}) error {
	var result error

	files := []string{".env", ".env.local"}
	for _, i := range cfg {
		if err := loadStruct(i, files...); err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result
}
