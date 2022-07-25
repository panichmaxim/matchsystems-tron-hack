package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/ethstore"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

var _ EthNeoService = (*serviceImpl)(nil)

type EthNeoService interface {
	ethstore.Store
}

func (s *serviceImpl) EthFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	return s.ethneo.EthFindAddressByHash(ctx, address)
}

func (s *serviceImpl) EthFindTransactionsByAddress(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	return s.ethneo.EthFindTransactionsByAddress(ctx, hash, page, pageSize)
}

func (s *serviceImpl) EthFindTransactionByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	return s.ethneo.EthFindTransactionByHash(ctx, hash)
}

func (s *serviceImpl) EthFindIncomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	return s.ethneo.EthFindIncomingTransactionAddress(ctx, hash)
}

func (s *serviceImpl) EthFindOutcomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	return s.ethneo.EthFindOutcomingTransactionAddress(ctx, hash)
}

func (s *serviceImpl) EthFindBlockByTransaction(ctx context.Context, hash string) (*neoutils.Node, error) {
	return s.ethneo.EthFindBlockByTransaction(ctx, hash)
}

func (s *serviceImpl) EthFindBlockByHeight(ctx context.Context, height string) (*neoutils.Node, error) {
	return s.ethneo.EthFindBlockByHeight(ctx, height)
}

func (s *serviceImpl) EthFindTransactionsInBlock(ctx context.Context, height string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	return s.ethneo.EthFindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *serviceImpl) EthFindAllInputAndOutputTransactions(ctx context.Context, hash string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	return s.ethneo.EthFindAllInputAndOutputTransactions(ctx, hash, page, pageSize)
}

func (s *serviceImpl) EthFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	return s.ethneo.EthFindBlockByHash(ctx, hash)
}

func (s *serviceImpl) EthFindMentionsByAddress(ctx context.Context, address string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	return s.ethneo.EthFindMentionsByAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) EthFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	return s.ethneo.EthFindContactByAddress(ctx, address)
}

func (s *serviceImpl) EthFindRiskScoreByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	return s.ethneo.EthFindRiskScoreByAddress(ctx, address)
}
