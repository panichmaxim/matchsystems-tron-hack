package btc

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

// BtcStore https://docs.google.com/document/d/1j1YvKeaVgDJR8-DFCOrTwvurX_pk056tLryNH2_1c80/edit#heading=h.2pk75mcpllz1
type BtcStore interface {
	neo4jstore.Store
	FindWalletByWid(ctx context.Context, wid string) (*neo4jstore.Node, error)
	FindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neo4jstore.Node, int, error)
	FindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neo4jstore.Node, int, error)
}
