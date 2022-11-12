package eth

import (
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

// EthStore https://docs.google.com/document/d/1Fl3D0NF2TO49oQDEFAOImCmeny5Nr4EZ/edit#
type EthStore interface {
	neo4jstore.Store
}
