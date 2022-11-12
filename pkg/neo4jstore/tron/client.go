package tron

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore/generic"
)

var _ TronStore = (*storeImpl)(nil)

type storeImpl struct {
	generic neo4jstore.Store
	session neo4j.SessionWithContext
}

func NewStore(session neo4j.SessionWithContext, category neo4jstore.CategoryRisk) TronStore {
	return &storeImpl{
		generic: generic.NewGenericStore(session, category),
		session: session,
	}
}
