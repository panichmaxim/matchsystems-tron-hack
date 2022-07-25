package app

import (
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/btcstore"
	"gitlab.com/rubin-dev/api/pkg/elastic"
	"gitlab.com/rubin-dev/api/pkg/ethstore"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"gitlab.com/rubin-dev/api/pkg/service/otlp"
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

	btcNeoClient, err := neoutils.CreateDriver(sharedConfig.BtcNeo.Address, sharedConfig.BtcNeo.Username, sharedConfig.BtcNeo.Password)
	if err != nil {
		return nil, errors.Wrap(err, "neoutils.CreateDriver")
	}

	ethNeoClient, err := neoutils.CreateDriver(sharedConfig.EthNeo.Address, sharedConfig.EthNeo.Username, sharedConfig.EthNeo.Password)
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

	btcneo := btcstore.NewStore(neoutils.CreateSession(btcNeoClient))
	ethneo := ethstore.NewStore(neoutils.CreateSession(ethNeoClient))
	svc := otlp.NewOTLPService(
		service.NewService(
			sqlImpl,
			jwtsvc,
			esClient,
			mailer.NewNotify(mailGate),
			btcneo,
			ethneo,
			sharedConfig.App.Dev,
		),
		app.tracer,
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
