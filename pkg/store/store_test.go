package store

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.com/rubin-dev/api/pkg/database"
	"os"
)

func loadFixtures(ctx context.Context, db *bun.DB, files []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return database.LoadFixtures(ctx, db, os.DirFS(cwd+"/fixtures"), files)
}
