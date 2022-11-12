package tron

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
	record1, err1 := s.FindAddressByHash(ctx, "41aeb466da22b8d280d7c57479c49ff061df311c74")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	// @todo
	//require.NotEmpty(t, record1.Props["risktag"])
	//require.NotEmpty(t, record1.Props["totalin"])

	record2, err2 := s.FindAddressByHash(ctx, "TRtxgyKrUXmjVks9GFyabPDu6K5952HWXV")
	require.NoError(t, err2)
	require.NotNil(t, record2)

	record3, err3 := s.FindAddressByHash(ctx, "TXcTPLwCERL3u7XMtxM9AAmujM1Ucx97eN")
	require.NoError(t, err3)
	require.NotNil(t, record3)

	// @todo
	//require.NotEmpty(t, record2.Props["risktag"])
	//require.NotEmpty(t, record2.Props["totalin"])
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
	m.CategoryFindByNameMock.Return(c, nil)
	s := NewStore(sess, m)
	record1, err1 := s.Risk(ctx, "TXcTPLwCERL3u7XMtxM9AAmujM1Ucx97eN")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	require.Nil(t, record1.Calculated)
	require.NotNil(t, record1.Reported)
	reported := *record1.Reported
	require.Equal(t, 111, reported.Category)
	require.Equal(t, float64(35), reported.Risk)
}

func TestStoreImpl_ConvertToHEXAddress(t *testing.T) {
	addr1, err1 := MustDecodeAddress("41a204f55f259535341e83080a945b28691130dfad")
	require.NoError(t, err1)
	require.Equal(t, "41a204f55f259535341e83080a945b28691130dfad", addr1)

	addr2, err2 := MustDecodeAddress("TQjtQf4JePBQadxLAjfSVhW2u632BHdmyp")
	require.NoError(t, err2)
	require.Equal(t, "41a204f55f259535341e83080a945b28691130dfad", addr2)
}

func TestStoreImpl_FindContactByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record1, err1 := s.FindContactByAddress(ctx, "TQjtQf4JePBQadxLAjfSVhW2u632BHdmyp")
	require.NoError(t, err1)
	require.NotNil(t, record1)

	record2, err2 := s.FindContactByAddress(ctx, "41a204f55f259535341e83080a945b28691130dfad")
	require.NoError(t, err2)
	require.NotNil(t, record2)
}

func TestStoreImpl_FindBlockByHash(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindBlockByHash(ctx, "000000000002ff1b8632a334b563960b0fe8ccf16a940308bac3b53165be7720")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindBlockByHeight(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindBlockByHeight(ctx, 196391)
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindBlockByTransaction(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)
	record, err := s.FindBlockByTransaction(ctx, "6078c24607224f4c692d11b84226ebbfb280e9b7b262f271918645a7edb46c80")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindTransactionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record1, total1, err1 := s.FindTransactionsByAddress(ctx, "TCuhmgyD2FkHV2vVcZzPg9uwTt9FgRZZUi", 1, 10)
	require.NoError(t, err1)
	require.NotNil(t, record1)
	require.Greater(t, total1, 0)

	record2, total2, err2 := s.FindTransactionsByAddress(ctx, "41203ebaa47ea3bdda77c1fa07bd542257d7dc35b2", 1, 10)
	require.NoError(t, err2)
	require.NotNil(t, record2)
	require.Greater(t, total2, 0)
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
	record, err := s.FindTransactionByHash(ctx, "6937ca320cec3fdb15e421da211994067d6766195e68dc7026a37597852e166b")
	require.NoError(t, err)
	require.NotNil(t, record)
}

func TestStoreImpl_FindMentionsByAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record1, total1, err1 := s.FindMentionsForAddress(ctx, "TXcTPLwCERL3u7XMtxM9AAmujM1Ucx97eN", 1, 10)
	require.NoError(t, err1)
	require.NotNil(t, record1)
	require.Greater(t, total1, 0)

	record2, total2, err2 := s.FindMentionsForAddress(ctx, "41a204f55f259535341e83080a945b28691130dfad", 1, 10)
	require.NoError(t, err2)
	require.NotNil(t, record2)
	require.Greater(t, total2, 0)
}

func TestStoreImpl_FindIncomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, total, err := s.FindIncomingTransactions(ctx, "61f0865630c01480a20399db848c4fb71c5a971521d0d873a4cd3887aafbe913", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Equal(t, total, 1)
}

func TestStoreImpl_FindOutcomingTransactionAddress(t *testing.T) {
	sess, err := createTestSession()
	require.NoError(t, err)
	require.NotNil(t, sess)

	s := NewStore(sess, nil)

	record, total, err := s.FindOutcomingTransactions(ctx, "61f0865630c01480a20399db848c4fb71c5a971521d0d873a4cd3887aafbe913", 1, 10)
	require.NoError(t, err)
	require.NotNil(t, record)
	require.Equal(t, total, 1)
}
