package ethstore

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/require"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"testing"
	"time"
)

const testNeoFirst = "neo4j://116.202.84.207"
const testNeoUsername = "neo4j"
const testNeoPassword = "RubinNeo#"

var ctx = context.TODO()

func createTestSession() (neo4j.SessionWithContext, error) {
	driver, err := createTestDriver()
	if err != nil {
		return nil, err
	}
	return neoutils.CreateSession(driver), nil
}

func createTestDriver() (neo4j.DriverWithContext, error) {
	return neoutils.CreateDriver(testNeoFirst, testNeoUsername, testNeoPassword)
}

func TestCreateDriver(t *testing.T) {
	driver, err := createTestDriver()
	require.NoError(t, err)
	require.NotNil(t, driver)
	timed, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	require.NoError(t, driver.VerifyConnectivity(timed))
}

func TestStoreImpl_EthFindAddressByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindAddressByHash(ctx, "0xe2211d98f0f89a9c5b61e39fc80fde9056d8e2c0")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindRiskScoreByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindRiskScoreByAddress(ctx, "0xaf933cf019e6f75a2d82f9bcb752dc6c1520ad76")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindContactByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindContactByAddress(ctx, "0xaf933cf019e6f75a2d82f9bcb752dc6c1520ad76")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindBlockByHash(ctx, "0x45877cf2f385d591320895fe080051471d2aac67534c0c2fc436cd663f6d00af")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindBlockByHeight(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindBlockByHeight(ctx, "1234")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindBlockByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindBlockByTransaction(ctx, "0x29b41d5f8a56b546e99bdc2f3449b3a703a6b86cb63d9e43f851b8cfd5a828eb")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindTransactionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, total, err := s.EthFindTransactionsByAddress(ctx, "0xea674fdde714fd979de3edf0f56aa9716b898ec8", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_EthFindAllInputAndOutputTransactions(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, total, err := s.EthFindAllInputAndOutputTransactions(ctx, "0x29b41d5f8a56b546e99bdc2f3449b3a703a6b86cb63d9e43f851b8cfd5a828eb", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_EthFindTransactionsInBlock(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, total, err := s.EthFindTransactionsInBlock(ctx, "12123746", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_EthFindTransactionByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindTransactionByHash(ctx, "0x29b41d5f8a56b546e99bdc2f3449b3a703a6b86cb63d9e43f851b8cfd5a828eb")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindMentionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, total, err := s.EthFindMentionsByAddress(ctx, "0xaf933cf019e6f75a2d82f9bcb752dc6c1520ad76", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Greater(t, total, 0)
}

func TestStoreImpl_EthFindIncomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindIncomingTransactionAddress(ctx, "0x5f21da012105b9deb306e13256922a0acb1460c50c255fdd56562a3151d7d96e")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_EthFindOutcomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess)
	record, err := s.EthFindOutcomingTransactionAddress(ctx, "0x5f21da012105b9deb306e13256922a0acb1460c50c255fdd56562a3151d7d96e")
	require.NoError(t, err)
	require.NotNil(t, record)
}
