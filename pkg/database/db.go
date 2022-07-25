package database

import (
	"context"
	"crypto/tls"
	"database/sql"
	"os"

	"github.com/uptrace/bun/extra/bunotel"
	"gitlab.com/falaleev-golang/zlog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// New вспомогательный метод для создания экземпляра bun.DB.
// Подробнее https://bun.uptrace.dev/guide/drivers.html
func New(ctx context.Context, dsn string, useTLS bool) (*bun.DB, error) {
	dbopts := []pgdriver.Option{
		pgdriver.WithDSN(dsn),
	}
	if useTLS {
		dbopts = append(dbopts, pgdriver.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(dbopts...))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(&zlog.QueryHook{})
	db.AddQueryHook(bunotel.NewQueryHook())

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

// NewTest вспомогательный метод для создания экземпляра bun.DB и
// создания схемы базы данных на лету.
func NewTest(ctx context.Context, dsn string, models []interface{}) (*bun.DB, error) {
	db, err := New(ctx, dsn, false)
	if err != nil {
		return nil, err
	}
	if err := CreateSchema(ctx, db, models); err != nil {
		return nil, err
	}
	return db, nil
}

func GetDefaultDsn() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if len(databaseURL) == 0 {
		databaseURL = "postgres://postgres:postgres@localhost:5432/app?sslmode=disable"
	}
	return databaseURL
}
