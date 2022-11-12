package tron

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

func (s *storeImpl) Health(ctx context.Context) error {
	return s.generic.Health(ctx)
}

func (s *storeImpl) Risk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	addr, err := MustEncodeAddress(address)
	if err != nil {
		return nil, err
	}

	return s.generic.Risk(ctx, addr)
}

func (s *storeImpl) FindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	addr, err := MustEncodeAddress(address)
	if err != nil {
		return nil, err
	}

	return s.generic.FindAddressByHash(ctx, addr)
}

func (s *storeImpl) FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	addr, err := MustEncodeAddress(address)
	if err != nil {
		return nil, 0, err
	}

	return s.generic.FindTransactionsByAddress(ctx, addr, page, pageSize)
}

func (s *storeImpl) FindTransactionByHash(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	return s.generic.FindTransactionByHash(ctx, txid)
}

func (s *storeImpl) FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindIncomingTransactions(ctx, txid, page, pageSize)
}

func (s *storeImpl) FindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindOutcomingTransactions(ctx, txid, page, pageSize)
}

func (s *storeImpl) FindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	return s.generic.FindBlockByTransaction(ctx, txid)
}

func (s *storeImpl) FindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error) {
	return s.generic.FindBlockByHeight(ctx, height)
}

func (s *storeImpl) FindTransactionsInBlock(ctx context.Context, height, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindTransactionsInBlock(ctx, height, page, pageSize)
}

func (s *storeImpl) FindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.generic.FindBlockByHash(ctx, hash)
}

func (s *storeImpl) FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	addr, err := MustEncodeAddress(address)
	if err != nil {
		return nil, 0, err
	}

	return s.generic.FindMentionsForAddress(ctx, addr, page, pageSize)
}

func (s *storeImpl) FindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	addr, err := MustEncodeAddress(address)
	if err != nil {
		return nil, errors.Wrap(err, "MustEncodeAddress")
	}

	return s.generic.FindContactByAddress(ctx, addr)
}

// FindWalletForAddress Сущность кошелек для адреса
func (s *storeImpl) FindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.generic.FindWalletForAddress(ctx, address)
}
