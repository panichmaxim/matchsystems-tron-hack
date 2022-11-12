package eth

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/models"
	"testing"
)

func TestStoreImpl_FindAddressByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindAddressByHash(ctx, "0xe2211d98f0f89a9c5b61e39fc80fde9056d8e2c0")
	require.NoError(t, err)
	require.NotNil(t, record)

	// @todo
	//require.NotEmpty(t, record.Props["risktag"])
	//require.NotEmpty(t, record.Props["totalin"])
}

func TestStoreImpl_Risk(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	m := NewCategoryRiskMock(minimock.NewController(t))
	m.CategoryFindByNumberMock.Return(&models.Category{
		Name:   "test",
		Number: 111,
		Risk:   100,
	}, nil)

	s := NewStore(sess, m)
	record1, err1 := s.Risk(ctx, "0xaf933cf019e6f75a2d82f9bcb752dc6c1520ad76")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.NotNil(t, record1.Reported)
	reported := *record1.Reported
	require.Equal(t, 111, reported.Category)
	require.Equal(t, float64(100), reported.Risk)
}

func TestStoreImpl_RiskCalculated(t *testing.T) {
	t.Skip("not supported yet")
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	m := NewCategoryRiskMock(minimock.NewController(t))
	m.CategoryFindByNumberMock.Return(nil, nil)

	s := NewStore(sess, m)
	record1, err1 := s.Risk(ctx, "0xea674fdde714fd979de3edf0f56aa9716b898ec8")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.Nil(t, record1.Calculated)
}

func TestStoreImpl_FindContactByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindContactByAddress(ctx, "0x34d5d28bff99082fd72d4a4f660a0c4b276a60ec")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindBlockByHash(ctx, "0x45877cf2f385d591320895fe080051471d2aac67534c0c2fc436cd663f6d00af")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindBlockByHeight(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindBlockByHeight(ctx, 1234)
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindBlockByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindBlockByTransaction(ctx, "0x29b41d5f8a56b546e99bdc2f3449b3a703a6b86cb63d9e43f851b8cfd5a828eb")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindTransactionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record1, total1, err := s.FindTransactionsByAddress(ctx, "0xea674fdde714fd979de3edf0f56aa9716b898ec8", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record1)
	require.Equal(t, total1, 43126689)

	record2, total2, err := s.FindTransactionsByAddress(ctx, "0xea674fdde714fd979de3edf0f56aa9716b898ec8", 2, 10)
	require.NoError(t, err)
	require.NotNil(t, record2)
	require.Equal(t, total2, 43126689)

	record3, total3, err := s.FindTransactionsByAddress(ctx, "0xea674fdde714fd979de3edf0f56aa9716b898ec8", 3, 10)
	require.NoError(t, err)
	require.NotNil(t, record3)
	require.Equal(t, total3, 43126689)

	require.NotEqual(t, record1[0].ID, record2[0].ID)
	require.NotEqual(t, record2[0].ID, record3[0].ID)
}

func TestStoreImpl_FindTransactionsInBlock(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, total, err := s.FindTransactionsInBlock(ctx, 12123746, 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_FindTransactionByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, err := s.FindTransactionByHash(ctx, "0x29b41d5f8a56b546e99bdc2f3449b3a703a6b86cb63d9e43f851b8cfd5a828eb")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindMentionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, total, err := s.FindMentionsForAddress(ctx, "0x34d5d28bff99082fd72d4a4f660a0c4b276a60ec", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_FindIncomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, total, err := s.FindIncomingTransactions(ctx, "0x5f21da012105b9deb306e13256922a0acb1460c50c255fdd56562a3151d7d96e", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Equal(t, total, 1)
}

func TestStoreImpl_FindOutcomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, total, err := s.FindOutcomingTransactions(ctx, "0x5f21da012105b9deb306e13256922a0acb1460c50c255fdd56562a3151d7d96e", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Equal(t, total, 1)
}
