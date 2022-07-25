package btcstore

import (
	"context"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

// FindAddressByHash Поиск адреса по хэшу.
//{
// 	"identity": 941113543,
// 	"labels": ["Address"],
// 	"properties": {
// 		"address": "1QHmaBDkDmKc7DpFtSRJrTHoPGGZwhEVeh",
// 		"type": "BTC",
// 		"category": "DDOS service"
// 	}
// }
func (s *storeImpl) BtcFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address: $address}) RETURN a`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "a")
}

// FindTransactionsByAddress Массив сущностей транзакций где участвует адрес
//{
//  "identity": 722254504,
//  "labels": [
//    "Transaction"
//  ],
//  "properties": {
//"txid": "ffaed7f59b96e59ea352be48a5281a426cb6fa7198929c9bcb4c43fefb105ec5"
//  }
//}
//
//{
//  "identity": 721894473,
//  "labels": [
//    "Transaction"
//  ],
//  "properties": {
//"txid": "f6262e91e05ae8d9fe054a9c1a4c7bf17127d3aa0494161c62443deed546249a"
//  }
//}
func (s *storeImpl) BtcFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	query := `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(t:Transaction) RETURN t SKIP $skip LIMIT $limit`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(t:Transaction) RETURN count(t) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"address": address,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindWalletAddresses Получение всех адресов кошелька по wid
// {
//  "identity": 913752974,
//  "labels": [
//    "Address"
//  ],
//  "properties": {
//"address": "1PrAybrWjRwnmcn9ytwhUJxqyB8HXYQ6fY"
//  }
//}
//
//{
//  "identity": 830258090,
//  "labels": [
//    "Address"
//  ],
//  "properties": {
//"address": "1EdrVvNVdmdNThnCztoCj4bVsiQY2Pqs1m"
//  }
//}
//
//...
//...
//
//{
//  "identity": 829971498,
//  "labels": [
//    "Address"
//  ],
//  "properties": {
//"address": "1EcUdQ23hhojfYmLk9cfMmiU5FJMHnhwX2"
//  }
//}
func (s *storeImpl) BtcFindWalletAddresses(ctx context.Context, wid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	wallet, err := s.BtcFindWalletByWid(ctx, wid)
	if err != nil {
		return nil, 0, err
	}
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	query := `MATCH (a:Address)-[:CS]->(w:Wallet {wid: $wid}) RETURN a SKIP $skip LIMIT $limit`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"wid":   wid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	var total int
	if v, ok := wallet.Props["addrcount"]; ok {
		total = cast.ToInt(v)
	}

	items, err := neoutils.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindWalletByAddress Получение информации о кошельке по wid
// {
//  "identity": 1110258087,
//  "labels": [
//    "Wallet"
//  ],
//  "properties": {
//"wid": "02d6a4d9aadb29b0c286450e55e1d0d2c2d65ed70889278e83079096b797bf80",
//"addrcount": 15,
//"type": "CS"
//  }
//}
func (s *storeImpl) BtcFindWalletByWid(ctx context.Context, wid string) (*neoutils.Node, error) {
	query := `MATCH (w:Wallet {wid: $wid}) RETURN w`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"wid": wid,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "w")
}

// FindWalletForAddress Сущность кошелек для адреса
//{
//  "identity": 1109449305,
//  "labels": [
//    "Wallet"
//  ],
//  "properties": {
//"wid": "a5aa688373f74b69de94602184aefe7a318e88a7950754162bba2a8ce24a4e9e",
//"addrcount": 52,
//"type": "CS"
//  }
//}
func (s *storeImpl) BtcFindWalletForAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address: $address}) WITH a MATCH (a)--(w:Wallet) RETURN w`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "w")
}

// FindContactByAddress Сущности contact для адреса
// {
//  "identity": 999007504,
//  "labels": [
//    "Contact"
//  ],
//  "properties": {
//"value": "bitcointalk.org"
//  }
//  }
func (s *storeImpl) BtcFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention)-[:INFO]->(c:Contact) RETURN c`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "c")
}

// FindMentionsForAddress Массив сущностей mention для адреса
//{
//  "identity": 999007702,
//  "labels": [
//    "Mention"
//  ],
//  "properties": {
//"date": "2022-04-21 17:48:47",
//"meta": "5261437445, userxl, 48597104-c181-11ec-823b-ecb1d77beb80, {"BTC": "1LDNLreKJ6GawBHPgB5yfVLBERi8g3SbQS"}, True, Smuggling drugs, США, английский, https://bitcointalk.org/index.php?topic=6460.msg94424#msg94424, Далее, английский, {"BTC": {"total": {"risk": 0.274}, "high": {"illegal": 0, "darkmarket": 0.057, "darkservice": 0, "fraud": 0.003, "gambling": 0.001, "mixer": 0.002, "ransom": 0, "scam": 0, "stolen": 0.001}, "moderate": {"atm": 0, "ml_high": 0.003, "ml_mod": 0.005, "ml_very_high": 0.105, "ptp_high": 0}, "min": {"market": 0, "miner": 0.101, "wallet": 0.683, "ptp_risk_low": 0.034, "payment": 0.002, "risk_low": 0.001}}}",
//"url": "https://bitcointalk.org/"
//  }
//}
func (s *storeImpl) BtcFindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN m SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (a:Address {address: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN count(m) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"address": address,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "m")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindRiskScore Получить risk_score для адреса
func (s *storeImpl) BtcFindRiskScore(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address: $address}) WITH a OPTIONAL MATCH (c:Category) WHERE c.category = a.category RETURN c`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "c")
}

// FindTransactionByHash Поиск транзакции по хэшу
//{
//  "identity": 442308,
//  "labels": [
//    "Transaction"
//  ],
//  "properties": {
//"nvout": 1,
//"txid": "c56cd9f576ed49fbdb8d2126b4ffdad956b061dd84000ef1c8bd218ffc5f56e3",
//"nvin": 2
//  }
//}
func (s *storeImpl) BtcFindTransactionByHash(ctx context.Context, txid string) (*neoutils.Node, error) {
	query := `MATCH (t:Transaction {txid: $txid}) RETURN t`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"txid": txid,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "t")
}

// FindIncomingTransactions Массив входящих адресов транзакции
// "18mddZvpUTLqb1twgokF5HtVbb45YJFhdB"
// "1LDNLreKJ6GawBHPgB5yfVLBERi8g3SbQS"
func (s *storeImpl) BtcFindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)-[:SENDS]->(t) RETURN a SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"txid":  txid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)-[:SENDS]->(t) RETURN count(a) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"txid": txid,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindBlockByHash Поиск блока по хэшу
// {
//  "identity": 940190120,
//  "labels": [
//    "Block"
//  ],
//  "properties": {
//    "time": "2022-04-01T17:28:49Z",
//    "hash": "000000000000000000077d011569f26c0ebb6deaad63a7fbbbf256badc61bbf6"
//  }
// }
func (s *storeImpl) BtcFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH (b:Block {hash: $hash}) RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

// FindTransactionsInBlockByHash Поиск транзакций в блоке по хэшу
func (s *storeImpl) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (b:Block {hash: $hash}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN t SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash":  hash,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (b:Block {hash: $hash}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN count(t) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"hash": hash,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindOutcomingTransactions Массив исходящих адресов транзакции
// "1LDNLreKJ6GawBHPgB5yfVLBERi8g3SbQS"
// "15pJdFf3EASU4NPTf5LS51LSAb8PEuvc2v"
func (s *storeImpl) BtcFindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)<-[:RECEIVES]-(t) RETURN a SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"txid":  txid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)<-[:RECEIVES]-(t) RETURN count(a) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"txid": txid,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindBlockByTransaction Сущность блок, в который входит транзакция
// {
//  "identity": 1696754437,
//  "labels": [
//    "Block"
//  ],
//  "properties": {
//"time": "2011-04-25T05:06:37Z",
//"hash": "000000000000951d92d4042ce87baf901341df61ea7435474477aefff4175a4d",
//"height": 120039
//  }
//}
func (s *storeImpl) BtcFindBlockByTransaction(ctx context.Context, txid string) (*neoutils.Node, error) {
	query := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (t)-[:BELONGS_TO]->(b:Block) RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"txid": txid,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

// FindBlockByNumber Поиск блока по номеру
// {
//  "identity": 1696754437,
//  "labels": [
//    "Block"
//  ],
//  "properties": {
//"time": "2011-04-25T05:06:37Z",
//"hash": "000000000000951d92d4042ce87baf901341df61ea7435474477aefff4175a4d",
//"height": 120039
//  }
//}
func (s *storeImpl) BtcFindBlockByNumber(ctx context.Context, height int) (*neoutils.Node, error) {
	query := `MATCH (b:Block {height: $height}) RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"height": height,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

// FindTransactionsInBlock Массив сущностей транзакций в блоке
// {
//  "identity": 431026,
//  "labels": [
//    "Transaction"
//  ],
//  "properties": {
//"nvout": 2,
//"txid": "36dd1c824e8583fbfde27dc4204a54cdc5614460afb579e9331dbf8cda790e75",
//"nvin": 1
//  }
//}
//
//{
//  "identity": 431025,
//  "labels": [
//    "Transaction"
//  ],
//  "properties": {
//"nvout": 2,
//"txid": "be1ddf3a3421665f7f662062bf64d9c122e2669b82b542ece28754ca54d9715a",
//"nvin": 1
//  }
//}
func (s *storeImpl) BtcFindTransactionsInBlock(ctx context.Context, height int, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (b:Block {height: $height}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN t SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"height": height,
		"skip":   skip,
		"limit":  limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (b:Block {height: $height}) WITH b MATCH (t:Transaction)-[:BELONGS_TO]->(b) RETURN count(t) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"height": height,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// FindAllInputAndOutputByTransaction Все инпуты/аутпуты у транзакции
// {
//  "identity": 882535748,
//  "labels": [
//    "Address"
//  ],
//  "properties": {
//"address": "1LAN1RjaNmfSC9qxgvcFUQo7pK41mp5Zsr"
//  }
//}
//
//{
//  "identity": 883468194,
//  "labels": [
//    "Address"
//  ],
//  "properties": {
//"address": "1LDNLreKJ6GawBHPgB5yfVLBERi8g3SbQS",
//"type": "BTC",
//"category": "Smuggling drugs"
//  }
//}
func (s *storeImpl) BtcFindAllInputAndOutputByTransaction(ctx context.Context, txid string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)--(t) RETURN a SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"txid":  txid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (t:Transaction {txid: $txid}) WITH t MATCH (a:Address)--(t) RETURN count(a) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"txid": txid,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
