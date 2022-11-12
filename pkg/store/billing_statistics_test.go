package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"os"
	"testing"
)

func TestStoreImpl_BillingStatisticsRiskRange(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{
		"billing_statistics_range.yaml",
	}))
	s := NewSQLStore(db)

	value, err := s.BillingStatisticsRiskRange(
		ctx,
		1,
		nil,
		nil,
		NetworkBtc,
		tools.Ptr[int](0),
		tools.Ptr[int](25),
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, 2, value)
}

func TestStoreImpl_BillingStatisticsRiskRanges(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{"billing_statistics_range.yaml"}))
	s := NewSQLStore(db)

	risks, err := s.BillingStatisticsRiskRanges(ctx, 1, nil, nil, NetworkBtc, nil)
	require.NoError(t, err)
	require.Equal(t, []int{2, 0, 0, 0}, risks)
}
