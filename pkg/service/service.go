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
	CategoryService
	BtcNeoService
	EthNeoService
	TronNeoService

	Health(ctx context.Context) error
	Close(ctx context.Context) error
}
