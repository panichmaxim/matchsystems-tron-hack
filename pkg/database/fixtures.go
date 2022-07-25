package database

import (
	"context"
	"io/fs"
	"text/template"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

// LoadFixtures вспомогательный метод для загрузки фикстур из переданной fs и списка
// yaml файлов. Поддерживается embed.FS.
func LoadFixtures(ctx context.Context, db *bun.DB, fsys fs.FS, files []string) error {
	funcMap := template.FuncMap{
		"now": func() string {
			return time.Now().Format(time.RFC3339Nano)
		},
		"next_day": func() string {
			return time.Now().Add(24 * time.Hour).Format(time.RFC3339Nano)
		},
		"next_week": func() string {
			return time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339Nano)
		},
	}
	return dbfixture.New(db, dbfixture.WithTemplateFuncs(funcMap)).Load(ctx, fsys, files...)
}
