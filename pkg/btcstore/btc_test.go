package btcstore

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindWalletByWid(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindWalletByWid(ctx, "02d6a4d9aadb29b0c286450e55e1d0d2c2d65ed70889278e83079096b797bf80")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindWalletAddresses(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, total, err := s.BtcFindWalletAddresses(ctx, "02d6a4d9aadb29b0c286450e55e1d0d2c2d65ed70889278e83079096b797bf80", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

//func TestFindWalletAddressesBroken(t *testing.T) {
//	sess, err := createTestSession()
//	require.NoError(t, err)
//	require.NotNil(t, sess)
//
//	s := NewStore(sess)
//	record, total, err := s.BtcFindWalletAddresses(ctx, "ee7edf4713c2dce4c99a81ad9963d46c5bd55053ca7188f34c5235418a0f58a9", 221850, 100)
//	require.NoError(t, err)
//	require.NotNil(t, record)
//	require.Greater(t, total, 0)
//}

func TestFindAddressByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindAddressByHash(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindBlockByHash(ctx, "000000000000000000077d011569f26c0ebb6deaad63a7fbbbf256badc61bbf6")
	require.NoError(t, err)
	require.NotNil(t, record)

	record, err = s.BtcFindBlockByHash(ctx, "000000002d9050318ec8112057423e30b9570b39998aacd00ca648216525fce3")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindTransactionsInBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, _, err := s.BtcFindTransactionsInBlockByHash(ctx, "000000002d9050318ec8112057423e30b9570b39998aacd00ca648216525fce3", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Len(t, record, 1)
}

func TestFindContactByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindContactByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindTransactionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindTransactionsByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 10)

	records2, _, err := s.BtcFindTransactionsByAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 2, 10)
	require.NoError(t, err)
	require.NotNil(t, records2)
	require.Len(t, records2, 10)

	require.NotEqual(t, records1[0].ID, records2[0].ID)
}

func TestFindWalletForAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindWalletForAddress(ctx, "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindMentionsForAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindMentionsForAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 1)

	records2, _, err := s.BtcFindMentionsForAddress(ctx, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", 2, 10)
	require.NoError(t, err)
	require.Nil(t, records2)
	require.Len(t, records2, 0)
}

func TestFindRiskScore(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindRiskScore(ctx, "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindTransactionByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.BtcFindTransactionByHash(ctx, "f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestFindIncomingTransactions(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindIncomingTransactions(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 10)
}

func TestFindOutcomingTransactions(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindOutcomingTransactions(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 2)
}

func TestFindBlockByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, err := s.BtcFindBlockByTransaction(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf")
	require.NoError(t, err)
	require.NotNil(t, records1)
}

func TestFindBlockByNumber(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, err := s.BtcFindBlockByNumber(ctx, 5000)
	require.NoError(t, err)
	require.NotNil(t, records1)
}

func TestFindTransactionsInBlock(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindTransactionsInBlock(ctx, 61234, 0, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 4)
}

func TestFindAllInputAndOutputByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	records1, _, err := s.BtcFindAllInputAndOutputByTransaction(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, records1)
	require.Len(t, records1, 10)

	records2, _, err := s.BtcFindAllInputAndOutputByTransaction(ctx, "677b67a894d2587c423976ed65131d5ea730d9bd164e7692beffc0441f40eebf", 2, 10)
	require.NoError(t, err)
	require.NotNil(t, records2)
	require.Len(t, records2, 6)
}
