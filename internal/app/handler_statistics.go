package app

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/store"
	"net/http"
)

func (app *App) statisticsHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get(ApiKeyHeaderKey)
	if len(apiKey) == 0 {
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "missing api token"})
		return
	}

	key, err := app.svc.BillingFindApiKey(r.Context(), apiKey)
	if err != nil {
		log.Err(err).Msg("app.svc.BillingFindApiKey")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}

	if key == nil {
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "invalid api token"})
		return
	}

	var from *string
	fromRaw := r.URL.Query().Get("from")
	if len(fromRaw) > 0 {
		from = &fromRaw
	}
	var to *string
	toRaw := r.URL.Query().Get("to")
	if len(toRaw) > 0 {
		to = &toRaw
	}
	fromTime, toTime := tools.ParseFromToDates(from, to)
	var duplicates *bool
	duplicatesRaw := r.URL.Query().Get("duplicates")
	if len(duplicatesRaw) > 0 {
		duplicates = tools.Ptr[bool](true)
	}

	btcStats, err := app.svc.BillingStatisticsNetwork(r.Context(), key.UserID, fromTime, toTime, store.NetworkBtc, duplicates)
	if err != nil {
		log.Err(err).Msg("app.svc.BillingFindApiKey")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}
	ethStats, err := app.svc.BillingStatisticsNetwork(r.Context(), key.UserID, fromTime, toTime, store.NetworkEth, duplicates)
	if err != nil {
		log.Err(err).Msg("app.svc.BillingFindApiKey")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}
	tronStats, err := app.svc.BillingStatisticsNetwork(r.Context(), key.UserID, fromTime, toTime, store.NetworkTron, duplicates)
	if err != nil {
		log.Err(err).Msg("app.svc.BillingFindApiKey")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}

	_ = json.NewEncoder(w).Encode(&models.BillingStatistics{
		Btc:  btcStats,
		Eth:  ethStats,
		Tron: tronStats,
	})
}
