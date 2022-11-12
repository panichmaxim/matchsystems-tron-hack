package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

var _ EthNeoService = (*serviceImpl)(nil)

type EthNeoService interface {
	EthHealth(ctx context.Context) error
	EthRisk(ctx context.Context, address string) (*neo4jstore.Risk, error)
	EthFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error)
	EthFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error)
	EthFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	EthFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	EthFindTransactionByHash(ctx context.Context, address string) (*neo4jstore.Node, error)
	EthFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	EthFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	EthFindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error)
	EthFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error)
	EthFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neo4jstore.Node, int, error)
	EthFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error)
}

func (s *serviceImpl) EthHealth(ctx context.Context) error {
	return s.ethneo.Health(ctx)
}

func (s *serviceImpl) EthRisk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	return s.ethneo.Risk(ctx, address)
}

func (s *serviceImpl) EthFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	return s.ethneo.FindAddressByHash(ctx, address)
}

func (s *serviceImpl) EthFindTransactionsByAddress(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.ethneo.FindTransactionsByAddress(ctx, hash, page, pageSize)
}

func (s *serviceImpl) EthFindTransactionByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.ethneo.FindTransactionByHash(ctx, hash)
}

func (s *serviceImpl) EthFindIncomingTransactions(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.ethneo.FindIncomingTransactions(ctx, hash, page, pageSize)
}

func (s *serviceImpl) EthFindOutcomingTransactions(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.ethneo.FindOutcomingTransactions(ctx, hash, page, pageSize)
}

func (s *serviceImpl) EthFindBlockByTransaction(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.ethneo.FindBlockByTransaction(ctx, hash)
}

func (s *serviceImpl) EthFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error) {
	return s.ethneo.FindBlockByHeight(ctx, height)
}

func (s *serviceImpl) EthFindTransactionsInBlock(ctx context.Context, height, page int, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.ethneo.FindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *serviceImpl) EthFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.ethneo.FindBlockByHash(ctx, hash)
}

func (s *serviceImpl) EthFindMentionsForAddress(ctx context.Context, address string, page int, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.ethneo.FindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) EthFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.ethneo.FindContactByAddress(ctx, address)
}
