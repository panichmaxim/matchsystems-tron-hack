package btc

import (
	"github.com/gojuno/minimock/v3"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/models"
	"testing"
)

func TestStoreImpl_FindWalletByWid(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	m := NewCategoryRiskMock(minimock.NewController(t))
	s := NewStore(sess, m)
	record, err := s.FindWalletByWid(ctx, "48b3f20a18818dd2261136608c6523a4146c6eb4ce349923b806f7e9db5436d6")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_Risk(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	m := NewCategoryRiskMock(minimock.NewController(t))
	c := &models.Category{
		Name:   "test",
		Number: 111,
		Risk:   35,
	}
	m.CategoryFindByNumberMock.Return(c, nil)
	s := NewStore(sess, m)
	record1, err1 := s.Risk(ctx, "111K8kZAEnJg245r2cM6y9zgJGHZtJPy6")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.NotNil(t, record1.Reported)
	reported := *record1.Reported
	require.Equal(t, 111, reported.Category)
	require.Equal(t, float64(35), reported.Risk)
}
func TestStoreImpl_RiskCalculated(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record1, err1 := s.Risk(ctx, "111sSKaK716T634qcYC9EgyGFdzwL2E8s")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.NotNil(t, record1.Calculated)
}

func TestStoreImpl_RiskWallet(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	mc := minimock.NewController(t)
	m := NewCategoryRiskMock(mc)
	m.CategoryFindByNumberMock.Return(&models.Category{
		Name:   "test",
		Number: 111,
		Risk:   35,
	}, nil)

	s := NewStore(sess, m)
	record1, err1 := s.Risk(ctx, "3AtZWUBbLwtLR3nTMUqXadSZSAQf9H962f")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.NotNil(t, record1.Wallet)
	reported := *record1.Wallet
	require.Equal(t, 111, reported.Category)
	require.Equal(t, float64(35), reported.Risk)
}

func TestStoreImpl_FindWalletAddresses(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, total, err := s.FindWalletAddresses(ctx, "02d6a4d9aadb29b0c286450e55e1d0d2c2d65ed70889278e83079096b797bf80", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_FindAddressByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindAddressByHash(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh")
	require.NoError(t, err)
	require.NotNil(t, record)

	require.NotEmpty(t, record.Total)
	require.NotEmpty(t, record.Address)
}

func TestStoreImpl_BtcFindBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	//record, err := s.FindBlockByHash(ctx, "000000000000000000077d011569f26c0ebb6deaad63a7fbbbf256badc61bbf6")
	//require.NoError(t, err)
	//require.NotNil(t, record)

	record, err := s.FindBlockByHash(ctx, "000000002d9050318ec8112057423e30b9570b39998aacd00ca648216525fce3")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindTransactionsInBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, _, err := s.FindTransactionsInBlockByHash(ctx, "000000002d9050318ec8112057423e30b9570b39998aacd00ca648216525fce3", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Len(t, record, 1)
}

func TestStoreImpl_FindContactByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindContactByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindTransactionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, _, err := s.FindTransactionsByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 10)

	records2, _, err := s.FindTransactionsByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 2, 10)
	require.NoError(t, err)
	require.NotNil(t, records2)
	require.Len(t, records2, 10)

	require.NotEqual(t, records1[0].ID, records2[0].ID)
}

func TestStoreImpl_FindWalletForAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindWalletForAddress(ctx, "14BYYSPnM4ve3NpRJj8k2idrQ7BJHN6w1G")
	require.NoError(t, err)
	require.NotNil(t, record)
	log.Debug().Msgf("%+v", record)
}

func TestStoreImpl_FindMentionsForAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, _, err := s.FindMentionsForAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 3)

	records2, _, err := s.FindMentionsForAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 2, 10)
	require.NoError(t, err)
	require.Nil(t, records2)
	require.Len(t, records2, 0)
}

func TestStoreImpl_FindTransactionByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindTransactionByHash(ctx, "f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindIncomingTransactions(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, count, err := s.FindIncomingTransactions(ctx, "7ecad9a3e67bebda52c0bdafb53863938b86e30e58eb179aec747903b50baa2d", 1, 10)
	require.NoError(t, err)
	require.Equal(t, count, 445)
	require.NotNil(t, records1)
	require.Len(t, records1, 10)
}

func TestStoreImpl_FindOutcomingTransactions(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, count, err := s.FindOutcomingTransactions(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf", 1, 10)
	require.NoError(t, err)
	require.Equal(t, count, 2)
	require.NotNil(t, records1)
	require.Len(t, records1, count)
}

func TestStoreImpl_FindBlockByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, err := s.FindBlockByTransaction(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf")
	require.NoError(t, err)
	require.NotNil(t, records1)
}

func TestStoreImpl_FindBlockByHeight(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, err := s.FindBlockByHeight(ctx, 5000)
	require.NoError(t, err)
	require.NotNil(t, records1)
}

func TestStoreImpl_FindTransactionsInBlock(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	records1, _, err := s.FindTransactionsInBlock(ctx, 61234, 0, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 4)
}
