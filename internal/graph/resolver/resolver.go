//go:generate go run github.com/99designs/gqlgen generate
package resolver

import (
	"gitlab.com/rubin-dev/api/pkg/mailer"
	"gitlab.com/rubin-dev/api/pkg/service"
)

func NewResolver(
	svc service.Service,
	notify mailer.Notify,
) *Resolver {
	return &Resolver{
		svc:    svc,
		notify: notify,
	}
}

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	notify mailer.Notify
	svc    service.Service
}
