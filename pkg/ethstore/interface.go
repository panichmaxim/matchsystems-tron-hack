package ethstore

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

type Store interface {
	// EthFindAddressByHash Поиск адреса по хэшу
	EthFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error)
	// EthFindTransactionsByAddress Массив сущностей транзакций где участвует адрес
	EthFindTransactionsByAddress(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error)
	// EthFindTransactionByHash Поиск транзакции по хэшу
	EthFindTransactionByHash(ctx context.Context, hash string) (*neoutils.Node, error)
	// EthFindIncomingTransactionAddress Входящий адрес транзакции
	EthFindIncomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error)
	// EthFindOutcomingTransactionAddress Исходящий адрес транзакции
	EthFindOutcomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error)
	// EthFindBlockByTransaction Сущность блок, в который входит транзакция
	EthFindBlockByTransaction(ctx context.Context, hash string) (*neoutils.Node, error)
	// EthFindBlockByHeight Поиск блока по номеру
	EthFindBlockByHeight(ctx context.Context, height string) (*neoutils.Node, error)
	// EthFindTransactionsInBlock Массив сущностей транзакций в блоке
	EthFindTransactionsInBlock(ctx context.Context, height string, page int, pageSize int) ([]*neoutils.Node, int, error)
	// EthFindAllInputAndOutputTransactions Инпут и аутпут транзакции
	EthFindAllInputAndOutputTransactions(ctx context.Context, hash string, page int, pageSize int) ([]*neoutils.Node, int, error)
	// EthFindBlockByHash Поиск блока по хэшу
	EthFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error)
	// EthFindMentionsByAddress Массив сущностей mention для адреса
	EthFindMentionsByAddress(ctx context.Context, address string, page int, pageSize int) ([]*neoutils.Node, int, error)
	// EthFindContactByAddress Сущности contact для адреса
	EthFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error)
	// EthFindRiskScoreByAddress Получить risk_score для адреса
	EthFindRiskScoreByAddress(ctx context.Context, address string) (*neoutils.Node, error)
}
