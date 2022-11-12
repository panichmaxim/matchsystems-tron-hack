package neo4jstore

import "context"

type Store interface {
	Health(ctx context.Context) error
	// FindWalletForAddress поиск кошелька по адресу
	FindWalletForAddress(ctx context.Context, address string) (*Node, error)
	// Risk reported и calculated риски
	Risk(ctx context.Context, address string) (*Risk, error)
	// FindContactByAddress Сущности contact для адреса
	FindContactByAddress(ctx context.Context, address string) (*Node, error)
	// FindAddressByHash Поиск адреса по хэшу
	FindAddressByHash(ctx context.Context, address string) (*FindAddressByHashNode, error)
	// FindTransactionsByAddress Массив сущностей транзакций, где участвует адрес
	FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*Node, int, error)
	// FindMentionsForAddress Массив сущностей mention для адреса
	FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*Node, int, error)
	// FindTransactionByHash Поиск транзакции по хэшу
	FindTransactionByHash(ctx context.Context, address string) (*Node, error)
	// FindIncomingTransactions Входящий адрес транзакции
	FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*Node, int, error)
	// FindOutcomingTransactions Исходящий адрес транзакции
	FindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*Node, int, error)
	// FindBlockByTransaction Сущность блок, в который входит транзакция
	FindBlockByTransaction(ctx context.Context, txid string) (*Node, error)
	// FindBlockByHeight Поиск блока по номеру
	FindBlockByHeight(ctx context.Context, height int) (*Node, error)
	// FindTransactionsInBlock Массив сущностей транзакций в блоке
	FindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*Node, int, error)
	// FindBlockByHash Поиск блока по хэшу
	FindBlockByHash(ctx context.Context, hash string) (*Node, error)
}
