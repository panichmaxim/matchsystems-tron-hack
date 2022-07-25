package zlog

import (
	"encoding/json"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"gitlab.com/falaleev-golang/zlog/stack"
)

func captureSentryException(event *sentry.Event) {
	sentryHub := sentry.CurrentHub()
	sentryHub.Client().CaptureEvent(event, nil, sentryHub.Scope())
}

func createSentryEvent(level zerolog.Level, p []byte) *sentry.Event {
	var erroredEvent ErrEvent
	_ = json.Unmarshal(p, &erroredEvent)

	return &sentry.Event{
		Level:     errorMap[level],
		Message:   erroredEvent.Error,
		Timestamp: time.Now(),
		Release:   erroredEvent.ShaCommit,
		Exception: convertToException(erroredEvent),
	}
}

type ErrEvent struct {
	Error     string        `json:"error"`
	ShaCommit string        `json:"sha"`
	Stack     []stack.Frame `json:"stack"`
}

var errorMap = map[zerolog.Level]sentry.Level{
	zerolog.DebugLevel: sentry.LevelDebug,
	zerolog.InfoLevel:  sentry.LevelInfo,
	zerolog.WarnLevel:  sentry.LevelWarning,
	zerolog.ErrorLevel: sentry.LevelError,
	zerolog.FatalLevel: sentry.LevelFatal,
	zerolog.PanicLevel: sentry.LevelFatal,
}

func convertFramesToSentry(frames []stack.Frame) (sf []sentry.Frame) {
	for i, f := range frames {
		if i == 0 {
			continue
		}

		sf = append(sf, sentry.Frame{
			Function: f.Function,
			Module:   f.Module,
			Filename: f.Filename,
			AbsPath:  f.Filename,
			Lineno:   f.Lineno,
			InApp:    true,
		})
	}

	for i, j := 0, len(sf)-1; i < j; i, j = i+1, j-1 {
		sf[i], sf[j] = sf[j], sf[i]
	}

	return
}
