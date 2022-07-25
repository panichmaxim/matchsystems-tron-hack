package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"testing"
)

func TestStoreImpl_BillingSmartRegisterRequest(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	s := NewSQLStore(db)

	_, err = s.BillingAddRequests(ctx, 1, 100)
	require.NoError(t, err)
	b, err := s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 100, b.Requests)

	_, err = s.BillingSmartRegisterRequest(ctx, 1, "test", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 99, b.Requests)

	_, err = s.BillingSmartRegisterRequest(ctx, 1, "test", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 99, b.Requests)

	_, err = s.BillingSmartRegisterRequest(ctx, 1, "foobar", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 98, b.Requests)
}

func TestStoreImpl_BillingRegisterRequest(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	s := NewSQLStore(db)

	_, err = s.BillingAddRequests(ctx, 1, 100)
	require.NoError(t, err)
	b, err := s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 100, b.Requests)

	_, err = s.BillingRegisterRequest(ctx, 1, "test", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 99, b.Requests)

	_, err = s.BillingRegisterRequest(ctx, 1, "test", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 98, b.Requests)

	_, err = s.BillingRegisterRequest(ctx, 1, "foobar", 1, "test", "test")
	require.NoError(t, err)
	b, err = s.BillingGetOrCreate(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 97, b.Requests)

	count, err := s.BillingHistoryRequestsCount(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 3, count)
}
