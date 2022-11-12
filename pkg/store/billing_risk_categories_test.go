package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"os"
	"testing"
)

func TestStoreImpl_BillingRiskFindCategories(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{
		"billing_find_categories.yaml",
		"category.yaml",
	}))
	s := NewSQLStore(db)

	risks, err := s.BillingRiskFindCategories(ctx, 1, nil, nil, NetworkBtc, nil)
	require.NoError(t, err)
	require.Len(t, risks, 2)
	require.Equal(t, 39, risks[0].ID)
	require.Equal(t, 2, risks[0].Count)
	require.Equal(t, "Corruption", risks[0].Name)
	require.Equal(t, 58, risks[1].ID)
	require.Equal(t, 1, risks[1].Count)
	require.Equal(t, "High Risk country", risks[1].Name)
}

func TestStoreImpl_BillingRiskFindDirectories(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{"billing_find_directories.yaml"}))
	s := NewSQLStore(db)

	risks, err := s.BillingRiskFindDirectories(ctx, 1, nil, nil, NetworkBtc, nil)
	require.NoError(t, err)
	require.Len(t, risks, 2)
	require.Equal(t, 1, risks[0].ID)
	require.Equal(t, 2, risks[0].Count)
	require.Equal(t, 2, risks[1].ID)
	require.Equal(t, 1, risks[1].Count)
}

func TestStoreImpl_BillingRiskFindCategoriesAndDirectories(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{
		"billing_find_categories.yaml",
		"billing_find_directories.yaml",
		"category.yaml",
	}))
	s := NewSQLStore(db)

	risk, err := s.BillingRiskFindCategoriesAndDirectories(
		ctx,
		1,
		nil,
		nil,
		NetworkBtc,
		nil,
	)
	require.NoError(t, err)
	require.NotNil(t, risk)
	require.Equal(t, map[string]int{
		"Corruption":                       2,
		"Drugs":                            1,
		"High Risk country":                1,
		"Stolen funds and Computer crimes": 2,
	}, risk)
}
