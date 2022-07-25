package generichttp

import (
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("test create", func(t *testing.T) {
		type Input struct {
			Name  string
			Value int
		}
		type Output struct {
			ID    string
			Name  string
			Value int
		}

		// test create function
		handler := Handler(func(r *http.Request, serializer Input) (object Output, err error) {
			return Output{}, nil
		})

		_ = handler

	})
}
