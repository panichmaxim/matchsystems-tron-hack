package btcstore

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

type BtcStore interface {
	BtcFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error)
	BtcFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error)
	BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindWalletForAddress(ctx context.Context, address string) (*neoutils.Node, error)
	BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindRiskScore(ctx context.Context, address string) (*neoutils.Node, error)
	BtcFindTransactionByHash(ctx context.Context, address string) (*neoutils.Node, error)
	BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindBlockByTransaction(ctx context.Context, txid string) (*neoutils.Node, error)
	BtcFindBlockByNumber(ctx context.Context, height int) (*neoutils.Node, error)
	BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error)
	BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error)
	BtcFindWalletByWid(ctx context.Context, wid string) (*neoutils.Node, error)
	BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neoutils.Node, int, error)
}

type Store interface {
	BtcStore
}

var _ Store = (*storeImpl)(nil)

type storeImpl struct {
	session neo4j.SessionWithContext
}

func NewStore(session neo4j.SessionWithContext) Store {
	return &storeImpl{session: session}
}

func (s *storeImpl) count(ctx context.Context, query string, values map[string]interface{}, key string) (int, error) {
	result, err := s.session.Run(ctx, query, values)
	if err != nil {
		return 0, err
	}
	if result.Next(ctx) {
		record := result.Record()
		item, ok := record.Get(key)
		if !ok {
			return 0, fmt.Errorf("invalid node: %v", record.Values)
		}

		return cast.ToInt(item), nil
	}

	return 0, result.Err()
}
