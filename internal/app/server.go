package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func (app *App) Close(ctx context.Context) error {
	return app.svc.Close(ctx)
}

func (app *App) createServer(ctx context.Context) *http.Server {
	return &http.Server{
		Handler:      app.router,
		Addr:         fmt.Sprintf(":%d", app.config.Server.ListenPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}
}

func (app *App) Listen(ctx context.Context) error {
	return app.createServer(ctx).ListenAndServe()
}
