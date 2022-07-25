package tpl

import (
	"context"
	"net/http"
)

type GlobalResolver = func(ctx context.Context) (interface{}, error)

type TemplateEngine interface {
	SetGlobalResolver(key string, resolver GlobalResolver)
	RenderWriter(ctx context.Context, w http.ResponseWriter, name string, data Data) error
	Render(ctx context.Context, name string, data Data) ([]byte, error)
	RenderError(w http.ResponseWriter, r *http.Request, statusCode int, data Data)
}
