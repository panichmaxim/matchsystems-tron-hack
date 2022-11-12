package tron

import (
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

type TronStore interface {
	neo4jstore.Store
}
