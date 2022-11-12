package btc

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/generic"
)

var _ BtcStore = (*storeImpl)(nil)

type storeImpl struct {
	session  neo4j.SessionWithContext
	generic  neo4jstore.Store
	category neo4jstore.CategoryRisk
}

func NewStore(
	session neo4j.SessionWithContext,
	category neo4jstore.CategoryRisk,
) BtcStore {
	return &storeImpl{
		generic:  generic.NewGenericStore(session, category),
		session:  session,
		category: category,
	}
}
