package app

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (app *App) categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)

	categories, _, err := app.svc.CategoryList(r.Context())
	if err != nil {
		log.Err(err).Msg("app.svc.CategoryList")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(categories)
	return
}
