package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

var _ BtcNeoService = (*serviceImpl)(nil)

type BtcNeoService interface {
	BtcHealth(ctx context.Context) error
	BtcRisk(ctx context.Context, address string) (*neo4jstore.Risk, error)
	BtcFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error)
	BtcFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error)
	BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindTransactionByHash(ctx context.Context, address string) (*neo4jstore.Node, error)
	BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error)
	BtcFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error)
	BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error)
	BtcFindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error)
	BtcFindWalletByWid(ctx context.Context, wid string) (*neo4jstore.Node, error)
	BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error)
}

func (s *serviceImpl) BtcHealth(ctx context.Context) error {
	return s.btcneo.Health(ctx)
}

func (s *serviceImpl) BtcFindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindContactByAddress(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.FindContactByAddress(ctx, address)
}

func (s *serviceImpl) BtcFindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	if err := s.v.BtcFindAddressByHash(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.FindAddressByHash(ctx, address)
}

func (s *serviceImpl) BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindTransactionsByAddress(ctx, address, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindTransactionsByAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) BtcFindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindWalletForAddress(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.FindWalletForAddress(ctx, address)
}

func (s *serviceImpl) BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindMentionsForAddress(ctx, address, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) BtcRisk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	if err := s.v.BtcFindRiskScore(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.Risk(ctx, address)
}

func (s *serviceImpl) BtcFindTransactionByHash(ctx context.Context, address string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindTransactionByHash(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.FindTransactionByHash(ctx, address)
}

func (s *serviceImpl) BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindIncomingTransactions(ctx, txid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindIncomingTransactions(ctx, txid, page, pageSize)
}

func (s *serviceImpl) BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindOutcomingTransactions(ctx, txid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindOutcomingTransactions(ctx, txid, page, pageSize)
}

func (s *serviceImpl) BtcFindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindBlockByTransaction(ctx, txid); err != nil {
		return nil, err
	}

	return s.btcneo.FindBlockByTransaction(ctx, txid)
}

func (s *serviceImpl) BtcFindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindBlockByNumber(ctx, height); err != nil {
		return nil, err
	}

	return s.btcneo.FindBlockByHeight(ctx, height)
}

func (s *serviceImpl) BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindTransactionsInBlock(ctx, height, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *serviceImpl) BtcFindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindBlockByHash(ctx, hash); err != nil {
		return nil, err
	}

	return s.btcneo.FindBlockByHash(ctx, hash)
}

func (s *serviceImpl) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindTransactionsInBlockByHash(ctx, hash, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindTransactionsInBlockByHash(ctx, hash, page, pageSize)
}

func (s *serviceImpl) BtcFindWalletByWid(ctx context.Context, wid string) (*neo4jstore.Node, error) {
	if err := s.v.BtcFindWalletByWid(ctx, wid); err != nil {
		return nil, err
	}

	return s.btcneo.FindWalletByWid(ctx, wid)
}

func (s *serviceImpl) BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	if err := s.v.BtcFindWalletAddresses(ctx, wid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.FindWalletAddresses(ctx, wid, page, pageSize)
}
