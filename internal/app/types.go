package app

import (
	"gitlab.com/rubin-dev/api/pkg/validator"
)

type ValidationErrorsResp struct {
	Errors validator.Errors `json:"errors"`
}
