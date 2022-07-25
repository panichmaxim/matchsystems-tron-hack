package tpl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/service"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path"
	"reflect"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

func NewJetTemplateEngine(
	svc service.Service,
	loader jet.Loader,
	isDev bool,
	router *mux.Router,
) (TemplateEngine, error) {
	var opts []jet.Option
	if isDev {
		opts = append(opts, jet.InDevelopmentMode())
	}

	views := jet.NewSet(loader, opts...)

	pwd, _ := os.Getwd()
	manifestPath := path.Join(pwd, "public/build/manifest.json")
	manifestBody, err := ioutil.ReadFile(manifestPath)
	manifest := map[string]string{}
	if err != nil {
		log.Err(err).Msg("manifest not found")
	} else {
		if err := json.Unmarshal(manifestBody, &manifest); err != nil {
			log.Err(err).Msg("error")
			return nil, err
		}
	}
	views.AddGlobal("isDev", isDev)
	views.AddGlobalFunc("asset", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("asset", 1, math.MaxInt16)
		asset := a.Get(0)
		if v, ok := manifest[asset.String()]; ok {
			return reflect.ValueOf(v)
		}

		return asset
	})
	views.AddGlobalFunc("hasPermission", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("hasPermission", 1, 10)

		u, ok := a.Get(0).Interface().(*models.User)
		if u == nil || !ok {
			return reflect.ValueOf(false)
		}

		var permissions []string
		for i := 1; i < a.NumOfArguments(); i++ {
			if a.Get(i).Type().String() == "int64" {
				permissions = append(permissions, cast.ToString(a.Get(i).Int()))
			} else {
				permissions = append(permissions, a.Get(i).String())
			}
		}

		return reflect.ValueOf(svc.UserHasPermissions(u, permissions))
	})

	views.AddGlobalFunc("path", func(a jet.Arguments) reflect.Value {
		a.RequireNumOfArguments("path", 1, 100)

		routeName := a.Get(0).String()
		if len(routeName) == 0 {
			log.Err(fmt.Errorf("missing or empty route name"))
		}

		var pairs []string
		for i := 1; i < a.NumOfArguments(); i++ {
			if a.Get(i).Type().String() == "int64" {
				pairs = append(pairs, cast.ToString(a.Get(i).Int()))
			} else {
				pairs = append(pairs, a.Get(i).String())
			}
		}

		url, err := router.Get(routeName).URL(pairs...)
		if err != nil {
			log.Err(err).Msg("error")
		}

		return reflect.ValueOf(url)
	})

	views.AddGlobal("is_dev", isDev)

	return &JetTemplateEngine{
		views:   views,
		globals: make(map[string]GlobalResolver),
	}, nil
}

type JetTemplateEngine struct {
	views   *jet.Set
	globals map[string]GlobalResolver
}

func (t *JetTemplateEngine) SetGlobalResolver(key string, resolver GlobalResolver) {
	t.globals[key] = resolver
}

func (t *JetTemplateEngine) render(ctx context.Context, w io.Writer, name string, data Data) error {
	if data == nil {
		data = Data{}
	}

	view, err := t.views.GetTemplate(name)
	if err != nil {
		log.Err(err).Msg("error")
		return err
	}

	if view == nil {
		return fmt.Errorf("view not found")
	}

	vars := make(jet.VarMap)
	for k := range data {
		vars.Set(k, data[k])
	}

	for key, resolver := range t.globals {
		v, err := resolver(ctx)
		if err != nil {
			log.Err(err).Msgf("global value resolver error in key %s = %s", key, err)
		}
		vars.Set(key, v)
	}

	return view.Execute(w, vars, nil)
}

func (t *JetTemplateEngine) Render(ctx context.Context, name string, data Data) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := t.render(ctx, buf, name, data); err != nil {
		log.Err(err).Msg("error")
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *JetTemplateEngine) RenderWriter(ctx context.Context, w http.ResponseWriter, name string, data Data) error {
	return t.render(ctx, w, name, data)
}
