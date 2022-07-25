package otlp

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
	"gitlab.com/rubin-dev/api/pkg/service"
)

var _ service.EthNeoService = (*metricService)(nil)

func (m *metricService) EthFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindAddressByHash")
	defer t.End()

	return m.s.EthFindAddressByHash(ctx, address)
}

func (m *metricService) EthFindTransactionsByAddress(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindTransactionsByAddress")
	defer t.End()

	return m.s.EthFindTransactionsByAddress(ctx, hash, page, pageSize)
}

func (m *metricService) EthFindTransactionByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindTransactionByHash")
	defer t.End()

	return m.s.EthFindTransactionByHash(ctx, hash)
}

func (m *metricService) EthFindIncomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindIncomingTransactionAddress")
	defer t.End()

	return m.s.EthFindIncomingTransactionAddress(ctx, hash)
}

func (m *metricService) EthFindOutcomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindOutcomingTransactionAddress")
	defer t.End()

	return m.s.EthFindOutcomingTransactionAddress(ctx, hash)
}

func (m *metricService) EthFindBlockByTransaction(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindBlockByTransaction")
	defer t.End()

	return m.s.EthFindBlockByTransaction(ctx, hash)
}

func (m *metricService) EthFindBlockByHeight(ctx context.Context, height string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindBlockByHeight")
	defer t.End()

	return m.s.EthFindBlockByHeight(ctx, height)
}

func (m *metricService) EthFindTransactionsInBlock(ctx context.Context, height string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindTransactionsInBlock")
	defer t.End()

	return m.s.EthFindTransactionsInBlock(ctx, height, page, pageSize)
}

func (m *metricService) EthFindAllInputAndOutputTransactions(ctx context.Context, hash string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindAllInputAndOutputTransactions")
	defer t.End()

	return m.s.EthFindAllInputAndOutputTransactions(ctx, hash, page, pageSize)
}

func (m *metricService) EthFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindBlockByHash")
	defer t.End()

	return m.s.EthFindBlockByHash(ctx, hash)
}

func (m *metricService) EthFindMentionsByAddress(ctx context.Context, address string, page int, pageSize int) ([]*neoutils.Node, int, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindMentionsByAddress")
	defer t.End()

	return m.s.EthFindMentionsByAddress(ctx, address, page, pageSize)
}

func (m *metricService) EthFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindContactByAddress")
	defer t.End()

	return m.s.EthFindContactByAddress(ctx, address)
}

func (m *metricService) EthFindRiskScoreByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	ctx, t := m.tracer.Start(ctx, "service.EthFindRiskScoreByAddress")
	defer t.End()

	return m.s.EthFindRiskScoreByAddress(ctx, address)
}
