package store

import (
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/database"
	"gitlab.com/rubin-dev/api/pkg/models"
	"os"
	"testing"
	"time"
)

func getDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, &time.Location{})
}

func TestStoreImpl_BillingHistoryList(t *testing.T) {
	ctx := context.TODO()
	db, err := database.NewTest(ctx, database.GetDefaultDsn(), models.GetModels())
	require.NoError(t, err)
	require.NoError(t, database.LoadFixtures(ctx, db, os.DirFS("./fixtures"), []string{"billing.yaml"}))
	s := NewSQLStore(db)

	now := time.Now().UTC()
	items, count, err := s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		nil,
		&now,
		tools.Ptr[bool](false),
	)
	require.NoError(t, err)
	require.Len(t, items, 4)
	require.Equal(t, 4, count)
	items, count, err = s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		nil,
		&now,
		nil,
	)
	require.NoError(t, err)
	require.Len(t, items, 4)
	require.Equal(t, 4, count)

	items, count, err = s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		nil,
		&now,
		tools.Ptr[bool](true),
	)
	require.NoError(t, err)
	require.Len(t, items, 3)
	require.Equal(t, 3, count)

	items, count, err = s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		&now,
		nil,
		tools.Ptr[bool](true),
	)
	require.NoError(t, err)
	require.Len(t, items, 3)
	require.Equal(t, 3, count)

	prevDay := getDate(now.Year(), now.Month(), now.Day()-1)
	items, count, err = s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		&prevDay,
		&now,
		tools.Ptr[bool](true),
	)
	require.NoError(t, err)
	require.Len(t, items, 2)
	require.Equal(t, 2, count)

	prevMonth := getDate(now.Year(), now.Month()-1, now.Day())
	nextMonth := getDate(now.Year(), now.Month()+1, now.Day())
	items, count, err = s.BillingHistoryList(
		ctx,
		1,
		100,
		0,
		&prevMonth,
		&nextMonth,
		tools.Ptr[bool](true),
	)
	require.NoError(t, err)
	require.Len(t, items, 5)
	require.Equal(t, 5, count)
}
