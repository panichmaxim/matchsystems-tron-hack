package service

import (
	"context"
)

type Service interface {
	AuthService
	UserService
	AccessRequestService
	ElasticService
	BillingService
	BtcNeoService
	EthNeoService
	CategoryService

	Health(ctx context.Context) error
	Close(ctx context.Context) error
}
