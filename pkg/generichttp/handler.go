package generichttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler converts function to http handler
// This function is generic over 2 parameters
// 	* S - serializer (request body json)
//	* O - object returned
func Handler[S any, O any](
	fn func(r *http.Request, serializer S) (object O, err error),
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" && r.Body == nil {
			writeError(w, http.StatusBadRequest, fmt.Errorf("empty body"))
			return
		}

		// unmarshal body into serializer
		into := new(S)
		if r.Method == "POST" {
			if err := json.NewDecoder(r.Body).Decode(into); err != nil {
				writeError(w, http.StatusInternalServerError, err)
				return
			}
		}

		result, err := fn(r, *into)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
	})
}
