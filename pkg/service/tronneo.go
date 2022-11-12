package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

var _ TronNeoService = (*serviceImpl)(nil)

type TronNeoService interface {
	TronHealth(ctx context.Context) error
	TronRisk(ctx context.Context, address string) (*neo4jstore.Risk, error)
	TronFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error)
	TronFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error)
	TronFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	TronFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	TronFindTransactionByHash(ctx context.Context, address string) (*neo4jstore.Node, error)
	TronFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	TronFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	TronFindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error)
	TronFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error)
	TronFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neo4jstore.Node, int, error)
	TronFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error)
}

func (s *serviceImpl) TronHealth(ctx context.Context) error {
	return s.tronneo.Health(ctx)
}

func (s *serviceImpl) TronRisk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	return s.tronneo.Risk(ctx, address)
}

func (s *serviceImpl) TronFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	return s.tronneo.FindAddressByHash(ctx, address)
}

func (s *serviceImpl) TronFindTransactionsByAddress(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.tronneo.FindTransactionsByAddress(ctx, hash, page, pageSize)
}

func (s *serviceImpl) TronFindTransactionByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.tronneo.FindTransactionByHash(ctx, hash)
}

func (s *serviceImpl) TronFindIncomingTransactions(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.tronneo.FindIncomingTransactions(ctx, hash, page, pageSize)
}

func (s *serviceImpl) TronFindOutcomingTransactions(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.tronneo.FindOutcomingTransactions(ctx, hash, page, pageSize)
}

func (s *serviceImpl) TronFindBlockByTransaction(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.tronneo.FindBlockByTransaction(ctx, hash)
}

func (s *serviceImpl) TronFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error) {
	return s.tronneo.FindBlockByHeight(ctx, height)
}

func (s *serviceImpl) TronFindTransactionsInBlock(ctx context.Context, height, page int, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.tronneo.FindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *serviceImpl) TronFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.tronneo.FindBlockByHash(ctx, hash)
}

func (s *serviceImpl) TronFindMentionsForAddress(ctx context.Context, address string, page int, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.tronneo.FindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) TronFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.tronneo.FindContactByAddress(ctx, address)
}
