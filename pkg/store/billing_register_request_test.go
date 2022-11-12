package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
	"testing"
)

func TestStoreImpl_BillingRegisterRequest(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	s := NewSQLStore(db)

	var uid int64 = 1

	_, err = s.BillingAddRequests(ctx, uid, 100)
	require.NoError(t, err)
	b, err := s.BillingGetOrCreate(ctx, uid)
	require.NoError(t, err)
	require.Equal(t, 100, b.Requests)

	r := &neo4jstore.Risk{
		Calculated: &neo4jstore.CalculatedRisk{
			Total:   1,
			Risk:    1,
			RiskRaw: 1,
			Items: map[int]*neo4jstore.CalculateItem{
				1: {Risk: 1, Total: 1, Percent: 1, ID: 1},
			},
		},
	}

	_, err = s.BillingRegisterRequest(ctx, uid, "test", r, NetworkBtc)
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, uid)
	require.NoError(t, err)
	require.Equal(t, 99, b.Requests)

	_, err = s.BillingRegisterRequest(ctx, uid, "test", r, NetworkBtc)
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, uid)
	require.NoError(t, err)
	require.Equal(t, 98, b.Requests)

	_, err = s.BillingRegisterRequest(ctx, uid, "foobar", r, NetworkBtc)
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, uid)
	require.NoError(t, err)
	require.Equal(t, 97, b.Requests)

	count, err := s.BillingHistoryRequestsCount(ctx, uid)
	require.NoError(t, err)
	require.Equal(t, 3, count)
}
