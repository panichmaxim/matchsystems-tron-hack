package validator

import (
	"gitlab.com/rubin-dev/api/pkg/store"
)

type Validation interface {
	UserValidation
	ElasticValidation
	BtcNeoValidation
}

func NewValidation(store store.Store) Validation {
	return &validationImpl{store}
}

type validationImpl struct {
	store store.Store
}
