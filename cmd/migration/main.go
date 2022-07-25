package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/falaleev-golang/zlog"
	"gitlab.com/rubin-dev/api/internal/cfg"
	"gitlab.com/rubin-dev/api/pkg/database"
	"os"
	"path/filepath"
)

func main() {
	sharedConfig, err := cfg.Load()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	lvl := zerolog.ErrorLevel
	if sharedConfig.App.Dev {
		lvl = zerolog.DebugLevel
	}
	log.Logger = zlog.Default(lvl, sharedConfig.Sentry.SentryDSN, "")

	target := os.Args[1]
	if len(target) == 0 {
		log.Fatal().Msg("target is missing")
	}

	destination, err := filepath.Abs(target)
	if err != nil {
		log.Fatal().Err(err).Msg("target directory for migration does not exist")
	}

	if err := database.CreateMigration(destination); err != nil {
		log.Fatal().Err(err).Msg("failed to create migration")
	}
}
