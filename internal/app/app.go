package app

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/elastic"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/btc"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/eth"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/tron"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/internal/cfg"
	"gitlab.com/rubin-dev/api/pkg/jwtoken"
	"gitlab.com/rubin-dev/api/pkg/mailer"
	"gitlab.com/rubin-dev/api/pkg/service"
	"gitlab.com/rubin-dev/api/pkg/store"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func NewApp(
	ctx context.Context,
	sharedConfig *cfg.DistributedConfig,
	sqlImpl store.Store,
) (*App, error) {
	mailGate, err := createMailGateway(sharedConfig.Mailgun, sharedConfig.Smtp, sharedConfig.App.SiteURL)
	if err != nil {
		return nil, errors.Wrap(err, "createMailGateway")
	}

	esClient, err := elastic.NewElastic(sharedConfig.Elastic.Addresses)
	if err != nil {
		return nil, errors.Wrap(err, "elastic.NewElastic")
	}

	btcNeoClient, err := neo4jstore.CreateDriver(ctx, sharedConfig.BtcNeo.Address, sharedConfig.BtcNeo.Username, sharedConfig.BtcNeo.Password)
	if err != nil {
		return nil, errors.Wrap(err, "neoutils.CreateDriver")
	}

	ethNeoClient, err := neo4jstore.CreateDriver(ctx, sharedConfig.EthNeo.Address, sharedConfig.EthNeo.Username, sharedConfig.EthNeo.Password)
	if err != nil {
		return nil, errors.Wrap(err, "neoutils.CreateDriver")
	}

	tronNeoClient, err := neo4jstore.CreateDriver(ctx, sharedConfig.TronNeo.Address, sharedConfig.TronNeo.Username, sharedConfig.TronNeo.Password)
	if err != nil {
		return nil, errors.Wrap(err, "neoutils.CreateDriver")
	}

	app := &App{
		config: sharedConfig,
		tracer: otel.Tracer("app-tracer"),
	}

	jwtsvc, err := jwtoken.NewService(
		[]byte(sharedConfig.Jwt.JwtSecret),
		time.Hour*1,
		time.Hour*24,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("jwt error")
	}

	btcneo := btc.NewStore(neo4jstore.CreateSession(btcNeoClient), sqlImpl)
	ethneo := eth.NewStore(neo4jstore.CreateSession(ethNeoClient), sqlImpl)
	tronneo := tron.NewStore(neo4jstore.CreateSession(tronNeoClient), sqlImpl)
	svc := service.NewServiceWithTracing(
		service.NewService(
			sqlImpl,
			jwtsvc,
			esClient,
			mailer.NewNotify(mailGate),
			btcneo,
			ethneo,
			tronneo,
			sharedConfig.App.Dev,
		),
		"service",
	)
	if err != nil {
		return nil, errors.Wrap(err, "otlp.NewOTLPService")
	}
	app.svc = svc

	app.router = app.createRouter(mux.NewRouter())

	return app, nil
}

type App struct {
	svc    service.Service
	router *mux.Router
	notify mailer.Notify
	tracer trace.Tracer
	config *cfg.DistributedConfig
}
