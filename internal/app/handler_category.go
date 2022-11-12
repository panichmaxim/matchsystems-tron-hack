package app

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/models"
	"net/http"
)

type CategoryGroup struct {
	CategoryGroup *models.CategoryGroup
	Categories    []*models.Category
}

func (app *App) categoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)

	var result []*CategoryGroup

	for _, categoryGroup := range app.svc.CategoryGroupList() {
		id := int64(categoryGroup.ID)
		categories, err := app.svc.CategoryList(r.Context(), &id)
		if err != nil {
			log.Err(err).Msg("app.svc.CategoryList")
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
			return
		}
		result = append(result, &CategoryGroup{
			CategoryGroup: categoryGroup,
			Categories:    categories,
		})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
	return
}
