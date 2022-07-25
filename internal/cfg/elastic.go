package cfg

import (
	"fmt"
	"strings"
)

type ElasticConfig struct {
	Addresses ElasticAddresses `env:"ELASTIC_ADDRESSES"`
}

type ElasticAddresses []string

func (s *ElasticAddresses) String() string {
	return fmt.Sprintf("%+v", *s)
}

func (s *ElasticAddresses) Set(value string) error {
	*s = strings.Split(value, ",")

	return nil
}
