package ethstore

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/spf13/cast"
	"gitlab.com/rubin-dev/api/pkg/neoutils"
)

type storeImpl struct {
	session neo4j.SessionWithContext
}

func NewStore(session neo4j.SessionWithContext) Store {
	return &storeImpl{session: session}
}

func (s *storeImpl) EthFindAddressByHash(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH(a:Address{address_string: $address}) RETURN a`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "a")
}

func (s *storeImpl) EthFindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	// query := `MATCH(inaddr:Address)-[t:TRANSACTION]-(outaddr:Address{address_string: $address}) WHERE id(t) > $cursor RETURN t ORDER BY id(t) LIMIT $limit`
	query := `MATCH(inaddr:Address)-[t:TRANSACTION]-(outaddr:Address{address_string: $address}) RETURN t SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH(a:Address {address_string: $address}) RETURN size((a)-[:TRANSACTION]-()) AS t`
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

func (s *storeImpl) EthFindTransactionByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH ()-[t:TRANSACTION{hash: $hash}]-() RETURN t`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "t")
}

func (s *storeImpl) EthFindIncomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH(inaddr:Address)-[t:TRANSACTION{hash: $hash}]->(outaddr:Address) RETURN inaddr`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "inaddr")
}

func (s *storeImpl) EthFindOutcomingTransactionAddress(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH(inaddr:Address)-[t:TRANSACTION{hash: $hash}]->(outaddr:Address) RETURN outaddr`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "outaddr")
}

func (s *storeImpl) EthFindBlockByTransaction(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH()-[t:TRANSACTION{hash: $hash}]-() MATCH(b:Block) WHERE b.number=t.block_number RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

func (s *storeImpl) EthFindBlockByHeight(ctx context.Context, height string) (*neoutils.Node, error) {
	query := `MATCH (b:Block {number: $height}) RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"height": height,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

func (s *storeImpl) EthFindTransactionsInBlock(ctx context.Context, height string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH()-[t:TRANSACTION]-() WHERE t.block_number = $height RETURN t SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"height": height,
		"skip":   skip,
		"limit":  limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH()-[t:TRANSACTION]-() WHERE t.block_number = $height RETURN count(t) AS t`
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

func (s *storeImpl) EthFindAllInputAndOutputTransactions(ctx context.Context, hash string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH(inaddr:Address)-[t:TRANSACTION{hash: $hash}]->(outaddr:Address) RETURN inaddr, outaddr SKIP $skip LIMIT $limit`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash":  hash,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH(inaddr:Address)-[t:TRANSACTION{hash: $hash}]->(outaddr:Address) RETURN count(t) AS t`
	total, err := s.count(ctx, countQuery, map[string]interface{}{
		"hash": hash,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neoutils.GetItems(ctx, result, "inaddr")
	if err != nil {
		return nil, 0, err
	}

	outcoming, err := neoutils.GetItems(ctx, result, "outaddr")
	if err != nil {
		return nil, 0, err
	}
	items = append(items, outcoming...)

	return items, total, nil
}

func (s *storeImpl) EthFindBlockByHash(ctx context.Context, hash string) (*neoutils.Node, error) {
	query := `MATCH(b:Block{hash: $hash}) RETURN b`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "b")
}

func (s *storeImpl) EthFindMentionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neoutils.Node, int, error) {
	query := `MATCH (a:Address {address_string: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN m`
	skip, limit := neoutils.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	countQuery := `MATCH (a:Address {address_string: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention) RETURN count(m) AS t`
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

func (s *storeImpl) EthFindContactByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address_string: $address}) WITH a MATCH (a)-[:FIGURED_IN]->(m:Mention)-[:INFO]->(c:Contact) RETURN c`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "c")
}

func (s *storeImpl) EthFindRiskScoreByAddress(ctx context.Context, address string) (*neoutils.Node, error) {
	query := `MATCH (a:Address {address_string: $address}) WITH a OPTIONAL MATCH (c:Category) WHERE c.category = a.category RETURN c`
	result, err := s.session.Run(ctx, query, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neoutils.GetItem(ctx, result, "c")
}

func (s *storeImpl) count(ctx context.Context, query string, values map[string]interface{}, key string) (int, error) {
	result, err := s.session.Run(ctx, query, values)
	if err != nil {
		return 0, err
	}
	if result.Next(ctx) {
		record := result.Record()
		item, ok := record.Get(key)
		if !ok {
			return 0, fmt.Errorf("invalid node: %v", record.Values)
		}

		return cast.ToInt(item), nil
	}

	return 0, result.Err()
}
