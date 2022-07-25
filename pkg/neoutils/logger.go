package neoutils

import (
	neolog "github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
	"github.com/rs/zerolog/log"
)

var _ neolog.Logger = (*neoLogger)(nil)

type neoLogger struct {
}

func (n *neoLogger) Error(name, id string, err error) {
	log.Error().Err(err).Str("name", name).Str("id", id).Msg("error")
}

func (n *neoLogger) Warnf(name string, id string, msg string, args ...interface{}) {
	log.Warn().Str("name", name).Str("id", id).Msgf(msg, args...)
}

func (n *neoLogger) Infof(name string, id string, msg string, args ...interface{}) {
	log.Info().Str("name", name).Str("id", id).Msgf(msg, args...)
}

func (n *neoLogger) Debugf(name string, id string, msg string, args ...interface{}) {
	log.Debug().Str("name", name).Str("id", id).Msgf(msg, args...)
}
