// nolint:wrapcheck
package zlog

import (
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var _ zerolog.LevelWriter = (*sentryWriter)(nil)

type sentryWriter struct {
	next zerolog.LevelWriter
}

func (s *sentryWriter) Write(p []byte) (n int, err error) {
	return s.next.Write(p)
}

func (s *sentryWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= zerolog.ErrorLevel {
		event := createSentryEvent(level, p)
		captureSentryException(event)
	}

	return s.next.WriteLevel(level, p)
}

func convertToException(e ErrEvent) []sentry.Exception {
	return []sentry.Exception{
		{
			Value: e.Error,
			Type:  e.Error,
			Stacktrace: &sentry.Stacktrace{
				Frames: convertFramesToSentry(e.Stack),
			},
		},
	}
}

// SentryWriter creates Plugin for zlog constructor.
// WARNING! sentry.Init use SENTRY_DSN environment variable if dsn argument is empty.
func SentryWriter(dsn string) Plugin {
	return func(next zerolog.LevelWriter) (zerolog.LevelWriter, error) {
		err := sentry.Init(sentry.ClientOptions{Dsn: dsn})

		return &sentryWriter{next: next}, errors.Wrap(err, "init sentry")
	}
}
