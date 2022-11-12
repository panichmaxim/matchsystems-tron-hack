package neo4jstore

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/log"
	"github.com/rs/zerolog"
)

var _ log.Logger = (*neoLogger)(nil)

func NewNeoLogger(logger *zerolog.Logger) *neoLogger {
	return &neoLogger{logger}
}

type neoLogger struct {
	logger *zerolog.Logger
}

func (n *neoLogger) Error(name, id string, err error) {
	n.logger.Error().Err(err).Str("name", name).Str("id", id).Msg("error")
}

func (n *neoLogger) Warnf(name string, id string, msg string, args ...interface{}) {
	n.logger.Warn().Str("name", name).Str("id", id).Msgf(msg, args...)
}

func (n *neoLogger) Infof(name string, id string, msg string, args ...interface{}) {
	n.logger.Info().Str("name", name).Str("id", id).Msgf(msg, args...)
}

func (n *neoLogger) Debugf(name string, id string, msg string, args ...interface{}) {
	n.logger.Debug().Str("name", name).Str("id", id).Msgf(msg, args...)
}
