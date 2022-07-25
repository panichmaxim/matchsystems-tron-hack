package generichttp

import (
	"encoding/json"
	"net/http"
)

type errorObject struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorObject{
		Error:      err.Error(),
		StatusCode: status,
	})
}
