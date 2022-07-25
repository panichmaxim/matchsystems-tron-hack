package app

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"gitlab.com/rubin-dev/api/pkg/store"
	"gitlab.com/rubin-dev/api/pkg/validator"
	"net/http"
)

const ApiKeyHeaderKey = "X-Api-Key"
const HeaderContentTypeKey = "Content-Type"
const HeaderContentTypeValue = "application/json"

type CheckResponse struct {
	Detail     string `json:"detail"`
	Address    string `json:"address"`
	Blockchain string `json:"blockchain"`
	Risk       int    `json:"risk,omitempty"`
	Category   string `json:"category"`
	Status     string `json:"status"`
	First      string `json:"first"`
	Last       string `json:"last"`
}

type RiskResponse struct {
	CheckResponse CheckResponse `json:"check_response"`
}

type ErrorResponse struct {
	Error string `json:"error"`
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

	var node *neoutils.Node
	var nodeErr error
	switch network {
	case "btc":
		node, nodeErr = app.svc.BtcFindRiskScore(r.Context(), address)
		break
	case "eth":
		node, nodeErr = app.svc.EthFindRiskScoreByAddress(r.Context(), address)
		break
	default:
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "unknown network"})
		return
	}

	if nodeErr != nil {
		log.Err(nodeErr).Msg("nodeErr")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)

		if errs, ok := nodeErr.(validator.Errors); ok {
			_ = json.NewEncoder(w).Encode(&errs)
			return
		}

		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}

	var risk int
	var category string

	if node != nil {
		if v, ok := node.Props["score"]; ok {
			risk = cast.ToInt(v)
		}
		if v, ok := node.Props["category"]; ok {
			category = cast.ToString(v)
		}

		if _, err := app.svc.BillingRegisterRequest(r.Context(), key.UserID, address, risk, category, network); err != nil {
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
		Category:   category,
	}

	if err := json.NewEncoder(w).Encode(&RiskResponse{CheckResponse: resp}); err != nil {
		log.Err(err).Msg("json.NewEncoder(w).Encode")
		w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&ErrorResponse{Error: "internal server error"})
		return
	}
}
