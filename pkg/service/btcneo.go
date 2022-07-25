package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/btcstore"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

var _ BtcNeoService = (*serviceImpl)(nil)

type BtcNeoService interface {
	btcstore.Store
}

func (s *serviceImpl) BtcFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	if err := s.v.FindContactByAddress(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindContactByAddress(ctx, address)
}

func (s *serviceImpl) BtcFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	if err := s.v.FindAddressByHash(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindAddressByHash(ctx, address)
}

func (s *serviceImpl) BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindTransactionsByAddress(ctx, address, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindTransactionsByAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) BtcFindWalletForAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	if err := s.v.FindWalletForAddress(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindWalletForAddress(ctx, address)
}

func (s *serviceImpl) BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindMentionsForAddress(ctx, address, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *serviceImpl) BtcFindRiskScore(ctx context.Context, address string) (*neoutils.Node, error) {
	if err := s.v.FindRiskScore(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindRiskScore(ctx, address)
}

func (s *serviceImpl) BtcFindTransactionByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	if err := s.v.FindTransactionByHash(ctx, address); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindTransactionByHash(ctx, address)
}

func (s *serviceImpl) BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindIncomingTransactions(ctx, txid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindIncomingTransactions(ctx, txid, page, pageSize)
}

func (s *serviceImpl) BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindOutcomingTransactions(ctx, txid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindOutcomingTransactions(ctx, txid, page, pageSize)
}

func (s *serviceImpl) BtcFindBlockByTransaction(ctx context.Context, txid string) (*neoutils.Node, error) {
	if err := s.v.FindBlockByTransaction(ctx, txid); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindBlockByTransaction(ctx, txid)
}

func (s *serviceImpl) BtcFindBlockByNumber(ctx context.Context, height int) (*neoutils.Node, error) {
	if err := s.v.FindBlockByNumber(ctx, height); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindBlockByNumber(ctx, height)
}

func (s *serviceImpl) BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindTransactionsInBlock(ctx, height, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *serviceImpl) BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindAllInputAndOutputByTransaction(ctx, txid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindAllInputAndOutputByTransaction(ctx, txid, page, pageSize)
}

func (s *serviceImpl) BtcFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	if err := s.v.FindBlockByHash(ctx, hash); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindBlockByHash(ctx, hash)
}

func (s *serviceImpl) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindTransactionsInBlockByHash(ctx, hash, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindTransactionsInBlockByHash(ctx, hash, page, pageSize)
}

func (s *serviceImpl) BtcFindWalletByWid(ctx context.Context, wid string) (*neoutils.Node, error) {
	if err := s.v.FindWalletByWid(ctx, wid); err != nil {
		return nil, err
	}

	return s.btcneo.BtcFindWalletByWid(ctx, wid)
}

func (s *serviceImpl) BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	if err := s.v.FindWalletAddresses(ctx, wid, page, pageSize); err != nil {
		return nil, 0, err
	}

	return s.btcneo.BtcFindWalletAddresses(ctx, wid, page, pageSize)
}
