package eth

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

func (s *storeImpl) Health(ctx context.Context) error {
	return s.generic.Health(ctx)
}

func (s *storeImpl) Risk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	return s.generic.Risk(ctx, address)
}

func (s *storeImpl) FindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	return s.generic.FindAddressByHash(ctx, address)
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
	return s.generic.FindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *storeImpl) FindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.generic.FindContactByAddress(ctx, address)
}

func (s *storeImpl) FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	query := `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[t:TRANSACTION]-(:Address) RETURN t SKIP $skip LIMIT $limit`
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH(a:Address {address: $address}) RETURN size((a)-[:TRANSACTION]-()) AS t`
	total, err := neo4jstore.Count(ctx, s.session, countQuery, map[string]interface{}{
		"address": address,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindWalletForAddress Сущность кошелек для адреса
func (s *storeImpl) FindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.generic.FindWalletForAddress(ctx, address)
}
