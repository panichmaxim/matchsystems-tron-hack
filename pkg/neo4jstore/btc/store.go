package btc

import (
	"context"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

func (s *storeImpl) Health(ctx context.Context) error {
	return s.generic.Health(ctx)
}

func (s *storeImpl) FindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	return s.generic.FindAddressByHash(ctx, address)
}

func (s *storeImpl) FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindTransactionsByAddress(ctx, address, page, pageSize)
}

func (s *storeImpl) FindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.generic.FindContactByAddress(ctx, address)
}

func (s *storeImpl) FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindMentionsForAddress(ctx, address, page, pageSize)
}

func (s *storeImpl) FindTransactionByHash(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	return s.generic.FindTransactionByHash(ctx, txid)
}

func (s *storeImpl) FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindIncomingTransactions(ctx, txid, page, pageSize)
}

func (s *storeImpl) FindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	return s.generic.FindBlockByHash(ctx, hash)
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

func (s *storeImpl) FindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	return s.generic.FindTransactionsInBlock(ctx, height, page, pageSize)
}

// FindWalletAddresses Получение всех адресов кошелька по wid
func (s *storeImpl) FindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	wallet, err := s.FindWalletByWid(ctx, wid)
	if err != nil {
		return nil, 0, err
	}

	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindWalletAddressesQuery, map[string]interface{}{
		"wid":   wid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var total int
	if v, ok := wallet.Props["addrcount"]; ok {
		total = cast.ToInt(v)
	}

	items, err := neo4jstore.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindWalletByWid Получение информации о кошельке по wid
func (s *storeImpl) FindWalletByWid(ctx context.Context, wid string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, FindWalletByWidQuery, map[string]interface{}{
		"wid": wid,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "w")
}

// FindWalletForAddress Сущность кошелек для адреса
func (s *storeImpl) FindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	return s.generic.FindWalletForAddress(ctx, address)
}

// FindTransactionsInBlockByHash Поиск транзакций в блоке по хэшу
func (s *storeImpl) FindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, FindTransactionsInBlockByHashQuery, map[string]interface{}{
		"hash":  hash,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, FindTransactionsInBlockByHashCountQuery, map[string]interface{}{
		"hash": hash,
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

func (s *storeImpl) Risk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	risk, err := s.generic.Risk(ctx, address)
	if err != nil {
		return nil, err
	}

	return risk, nil
}
