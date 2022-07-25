package app

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gitlab.com/falaleev-golang/zlog"
	"gitlab.com/rubin-dev/api/internal/graph/dataloader"
	"gitlab.com/rubin-dev/api/internal/graph/extensions"
	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/internal/graph/resolver"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"net/http"
	"time"
)

func (app *App) createRouter(router *mux.Router) *mux.Router {
	router.Use(otelmux.Middleware("app"))
	router.Use(zlog.HTTPAccessLogMiddleware)
	router.Use(zlog.HTTPRecoverMiddleware)

	app.createGraphqlRouter(router)

	router.HandleFunc("/api/v1/{network}/risk/{address}", app.riskHandler)
	router.HandleFunc("/api/v1/category", app.categoryHandler)

	return router
}

func (app *App) createGraphqlRouter(r *mux.Router) {
	ad := resolver.NewAuthDirective(app.svc, app.tracer)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Directives: generated.DirectiveRoot{
			Auth:        ad.Auth,
			AuthRefresh: ad.AuthRefresh,
		},
		Resolvers: resolver.NewResolver(
			app.svc,
			app.notify,
		),
	}))

	srv.Use(extensions.NewOperationLoggerExtension())
	srv.Use(extensions.NewOperationTracerExtension(app.tracer))
	//srv.Use(extensions.NewFieldTracerExtension(app.tracer))

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		if v, ok := err.(*gqlerror.Error); ok {
			if v.Message != resolver.AccessDenied {
				log.Err(err).Msg("graphql err")
			}
		}
		return graphql.DefaultErrorPresenter(ctx, err)
	})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Err(err.(error)).Msg("graphql fatal error")
		return gqlerror.Errorf("Internal server error!")
	})

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return ctx, nil
		},
	})
	srv.Use(extension.Introspection{})

	r.Handle("/graphql/play", playground.Handler("GraphQL playground", "/graphql"))

	middleware := dataloader.Middleware(app.svc, srv)
	corsImpl := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodHead, http.MethodGet, http.MethodPost},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	r.Handle("/graphql", jwtoken.BearerAuthorizationMiddleware(corsImpl.Handler(middleware)))
}
