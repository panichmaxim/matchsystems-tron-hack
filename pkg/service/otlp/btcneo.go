package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.BtcNeoService = (*metricService)(nil)

func (m *metricService) BtcFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindContactByAddress")
	defer t.End()

	return m.s.BtcFindContactByAddress(ctx, address)
}

func (m *metricService) BtcFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindAddressByHash")
	defer t.End()

	return m.s.BtcFindAddressByHash(ctx, address)
}

func (m *metricService) BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindTransactionsByAddress")
	defer t.End()

	return m.s.BtcFindTransactionsByAddress(ctx, address, page, pageSize)
}

func (m *metricService) BtcFindWalletForAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindWalletForAddress")
	defer t.End()

	return m.s.BtcFindWalletForAddress(ctx, address)
}

func (m *metricService) BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindMentionsForAddress")
	defer t.End()

	return m.s.BtcFindMentionsForAddress(ctx, address, page, pageSize)
}

func (m *metricService) BtcFindRiskScore(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindRiskScore")
	defer t.End()

	return m.s.BtcFindRiskScore(ctx, address)
}

func (m *metricService) BtcFindTransactionByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindTransactionByHash")
	defer t.End()

	return m.s.BtcFindTransactionByHash(ctx, address)
}

func (m *metricService) BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindIncomingTransactions")
	defer t.End()

	return m.s.BtcFindIncomingTransactions(ctx, txid, page, pageSize)
}

func (m *metricService) BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindContactByAddress")
	defer t.End()

	return m.s.BtcFindOutcomingTransactions(ctx, txid, page, pageSize)
}

func (m *metricService) BtcFindBlockByTransaction(ctx context.Context, txid string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindBlockByTransaction")
	defer t.End()

	return m.s.BtcFindBlockByTransaction(ctx, txid)
}

func (m *metricService) BtcFindBlockByNumber(ctx context.Context, height int) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindBlockByNumber")
	defer t.End()

	return m.s.BtcFindBlockByNumber(ctx, height)
}

func (m *metricService) BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindTransactionsInBlock")
	defer t.End()

	return m.s.BtcFindTransactionsInBlock(ctx, height, page, pageSize)
}

func (m *metricService) BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindAllInputAndOutputByTransaction")
	defer t.End()

	return m.s.BtcFindAllInputAndOutputByTransaction(ctx, txid, page, pageSize)
}

func (m *metricService) BtcFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindBlockByHash")
	defer t.End()

	return m.s.BtcFindBlockByHash(ctx, hash)
}

func (m *metricService) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindTransactionsInBlockByHash")
	defer t.End()

	return m.s.BtcFindTransactionsInBlockByHash(ctx, hash, page, pageSize)
}

func (m *metricService) BtcFindWalletByWid(ctx context.Context, wid string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindWalletByWid")
	defer t.End()

	return m.s.BtcFindWalletByWid(ctx, wid)
}

func (m *metricService) BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.BtcFindWalletAddresses")
	defer t.End()

	return m.s.BtcFindWalletAddresses(ctx, wid, page, pageSize)
}
