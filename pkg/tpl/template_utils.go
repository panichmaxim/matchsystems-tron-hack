package tpl

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (t *JetTemplateEngine) RenderError(w http.ResponseWriter, r *http.Request, statusCode int, data Data) {
	w.WriteHeader(statusCode)
	if data == nil {
		data = Data{}
	}
	data.Set("statusCode", statusCode)
	err := t.RenderWriter(r.Context(), w, "error.jet.html", data)
	if err != nil {
		log.Err(err).Msg("error while RenderTpl error.jet.html template")
		return
	}
}
