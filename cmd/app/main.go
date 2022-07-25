package main

import (
	"context"
	"embed"
	"os"
	"time"

	"gitlab.com/falaleev-golang/zlog"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/uptrace-go/uptrace"
	"gitlab.com/rubin-dev/api/internal/app"
	"gitlab.com/rubin-dev/api/internal/cfg"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/store"
)

var shaCommit = "local"

//go:embed migrations/*
var migrationFS embed.FS

//go:embed fixtures/*
var fixturesFS embed.FS

var fixtures = []string{
	"fixtures/base.yaml",
}

func main() {
	sharedConfig, err := cfg.Load()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if sharedConfig.App.Dev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		log.Logger = zlog.Default(zerolog.ErrorLevel, sharedConfig.Sentry.SentryDSN, shaCommit)
	}
	ctx := log.Logger.WithContext(context.Background())

	if len(sharedConfig.Uptrace.UptraceDSN) > 0 {
		uptrace.ConfigureOpentelemetry(
			uptrace.WithDSN(sharedConfig.Uptrace.UptraceDSN),
			uptrace.WithServiceName(sharedConfig.Uptrace.UptraceName),
			uptrace.WithServiceVersion(shaCommit),
		)

		// SendText buffered spans and free resources.
		defer uptrace.Shutdown(ctx)
	}

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	db, err := database.New(ctx, sharedConfig.Database.DatabaseURL, !sharedConfig.App.Dev)
	if err != nil {
		log.Fatal().Err(err).Msg("error")
	}
	defer func(db *bun.DB) {
		if err := db.Close(); err != nil {
			log.Err(err).Msg("db close error")
		}
	}(db)

	sqlImpl := store.NewSQLStore(db)
	if sharedConfig.Database.LoadMigrations {
		if err := sqlImpl.LoadMigrations(ctx, migrationFS); err != nil {
			log.Fatal().Err(err).Msg("error")
		}
		return
	}
	if sharedConfig.Database.LoadFixtures {
		if err := sqlImpl.LoadFixtures(ctx, fixturesFS, fixtures); err != nil {
			log.Fatal().Err(err).Msg("error")
		}
	}

	application, err := app.NewApp(sharedConfig, sqlImpl)
	if err != nil {
		log.Fatal().Err(err).Msg("err")
	}
	defer func(application *app.App, ctx context.Context) {
		err := application.Close(ctx)
		if err != nil {
			log.Err(err).Msg("application close error")
		}
	}(application, ctx)

	log.Info().Msgf("listen :%d", sharedConfig.Server.ListenPort)
	log.Fatal().Err(application.Listen(ctx)).Msg("failed to start server")
}
