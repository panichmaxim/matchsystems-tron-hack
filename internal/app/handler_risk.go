package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"gitlab.com/rubin-dev/api/pkg/store"
	"gitlab.com/rubin-dev/api/pkg/validator"
	"net/http"
)

const ApiKeyHeaderKey = "X-Api-Key"
const HeaderContentTypeKey = "Content-Type"
const HeaderContentTypeValue = "application/json"

type CheckResponse struct {
	Detail     string           `json:"detail"`
	Address    string           `json:"address"`
	Blockchain string           `json:"blockchain"`
	Risk       *neo4jstore.Risk `json:"risk,omitempty"`
	Category   string           `json:"category"`
	Status     string           `json:"status"`
	First      string           `json:"first"`
	Last       string           `json:"last"`
}

type RiskResponse struct {
	CheckResponse CheckResponse `json:"check_response"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (app *App) getRisk(ctx context.Context, network, address string) (*neo4jstore.Risk, error) {
	switch network {
	case "btc":
		return app.svc.BtcRisk(ctx, address)

	case "eth":
		return app.svc.EthRisk(ctx, address)

	case "tron":
		return app.svc.TronRisk(ctx, address)
	}

	return nil, fmt.Errorf("unknown network")
}

func (app *App) riskHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	address := vars["address"]
	if len(address) == 0 {
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "empty address"})
		return
	}

	network := vars["network"]
	if len(network) == 0 {
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "empty network"})
		return
	}

	risk, err := app.getRisk(r.Context(), network, address)
	if err != nil {
		log.Err(err).Msg("riskErr")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)

		if errs, ok := err.(validator.Errors); ok {
			_ = json.NewEncoder(w).Encode(&errs)
			return
		}

		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}

	if risk != nil {
		if _, err := app.svc.BillingRegisterRequest(r.Context(), key.UserID, address, risk, network); err != nil {
			log.Err(err).Msg("app.svc.BillingRegisterRequest")
			w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)

			if errors.Is(err, store.ErrInsufficientBalance) {
				w.WriteHeader(http.StatusPaymentRequired)
				_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "insufficient balance"})
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
			return
		}
	}

	resp := CheckResponse{
		Detail:     "Success",
		Address:    address,
		Blockchain: network,
		Risk:       risk,
	}

	if err := json.NewEncoder(w).Encode(&RiskResponse{CheckResponse: resp}); err != nil {
		log.Err(err).Msg("json.NewEncoder(w).Encode")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}
}
