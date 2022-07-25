package zlog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Plugin func(zerolog.LevelWriter) (zerolog.LevelWriter, error)

// Default create new default preconfigured logger with sentry integration and commit sha field.
func Default(lvl zerolog.Level, sentryDSN, commitSHA string) zerolog.Logger {
	zerolog.ErrorStackMarshaler = MarshalStack
	return New(os.Stdout, SentryWriter(sentryDSN)).
		Level(lvl).
		With().
		Stack().
		Caller().
		Str("sha", commitSHA).
		Logger()
}

// New creates zerolog instance with zerolog.LevelWriter middlewares.
func New(w io.Writer, plugins ...Plugin) zerolog.Logger {
	var levelWriter zerolog.LevelWriter = levelWriterAdapter{w}

	errors := make([]error, 0)

	for _, p := range plugins {
		if writer, err := p(levelWriter); err != nil {
			errors = append(errors, err)
		} else {
			levelWriter = writer
		}
	}

	l := zerolog.New(levelWriter)
	for _, err := range errors {
		l.Err(err).Send()
	}

	return l
}

type levelWriterAdapter struct {
	io.Writer
}

// nolint:wrapcheck
func (a levelWriterAdapter) WriteLevel(_ zerolog.Level, p []byte) (int, error) {
	return a.Write(p)
}
