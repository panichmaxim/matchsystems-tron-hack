package app

import (
	"encoding/json"
	"net/http"
)

func (app *App) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)

	if err := app.svc.Health(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]any{"error": err})
	}

	w.WriteHeader(http.StatusOK)
	return
}
