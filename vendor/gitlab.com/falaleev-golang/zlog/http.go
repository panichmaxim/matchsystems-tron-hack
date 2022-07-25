package zlog

import (
	"encoding/json"
	"fmt"
	"gitlab.com/falaleev-golang/zlog/stack"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

// HTTPAccessLogMiddleware won't work if you forget to set base or connection context with zerolog instance.
// You can see an example in the test.
func HTTPAccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		next.ServeHTTP(ww, r)

		zerolog.Ctx(r.Context()).Info().
			Str("type", "access").
			Timestamp().
			Fields(map[string]interface{}{
				"remote_ip":  r.RemoteAddr,
				"url":        r.URL.Path,
				"proto":      r.Proto,
				"method":     r.Method,
				"user_agent": r.Header.Get("User-Agent"),
				"status":     ww.Status(),
				"latency":    time.Since(start).String(),
				"bytes_in":   r.Header.Get("Content-Length"),
				"bytes_out":  ww.BytesWritten(),
			}).
			Send()
	})
}

type httpErr struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Stack   []stack.Frame `json:"stack"`
}

// HTTPRecoverMiddleware won't work if you forget to set base or connection context with zerolog instance.
// You can see an example in the test.
func HTTPRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				//nolint:goerr113
				e := fmt.Errorf("%s", err)

				zerolog.Ctx(r.Context()).
					Error().
					Err(e).
					Str("type", "panic").
					Timestamp().
					Send()

				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(httpErr{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Stack:   stack.GetStacktrace(stackSkip),
				})

				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}
