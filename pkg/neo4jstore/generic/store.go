package generic

import (
	"context"
	"github.com/spf13/cast"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

func NewGenericStore(session neo4j.SessionWithContext, category neo4jstore.CategoryRisk) neo4jstore.Store {
	return &GenericStoreImpl{session, category}
}

type GenericStoreImpl struct {
	session  neo4j.SessionWithContext
	category neo4jstore.CategoryRisk
}

func (s *GenericStoreImpl) Health(ctx context.Context) error {
	_, err := s.session.Run(ctx, `return 1`, map[string]any{})
	return err
}

func (s *GenericStoreImpl) FindAddressByHash(ctx context.Context, address string) (*neo4jstore.FindAddressByHashNode, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindAddressByHashQuery, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "result.Single")
	}

	return neo4jstore.ConvertType[neo4jstore.FindAddressByHashNode](record.Values[0])
}

func (s *GenericStoreImpl) FindTransactionsByAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindTransactionsByAddressQuery, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, neo4jstore.FindTransactionsByAddressCountQuery, map[string]interface{}{
		"address": address,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *GenericStoreImpl) FindTransactionByHash(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindTransactionByHashQuery, map[string]interface{}{
		"txid": txid,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "t")
}

func (s *GenericStoreImpl) FindIncomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindIncomingTransactionsQuery, map[string]interface{}{
		"txid":  txid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, neo4jstore.FindIncomingTransactionsCountQuery, map[string]interface{}{
		"txid": txid,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *GenericStoreImpl) FindOutcomingTransactions(ctx context.Context, txid string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindOutcomingTransactionsQuery, map[string]interface{}{
		"txid":  txid,
		"skip":  skip,
		"limit": limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, neo4jstore.FindOutcomingTransactionsCountQuery, map[string]interface{}{
		"txid": txid,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "a")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *GenericStoreImpl) FindBlockByTransaction(ctx context.Context, txid string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindBlockByTransactionQuery, map[string]interface{}{
		"txid": txid,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "b")
}

func (s *GenericStoreImpl) FindBlockByHeight(ctx context.Context, height int) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindBlockByHeightQuery, map[string]interface{}{
		"height": height,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "b")
}

func (s *GenericStoreImpl) FindTransactionsInBlock(ctx context.Context, height, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindTransactionsInBlockQuery, map[string]interface{}{
		"height": height,
		"skip":   skip,
		"limit":  limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, neo4jstore.FindTransactionsInBlockCountQuery, map[string]interface{}{
		"height": height,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "t")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *GenericStoreImpl) FindBlockByHash(ctx context.Context, hash string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindBlockByHashQuery, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "b")
}

func (s *GenericStoreImpl) FindMentionsForAddress(ctx context.Context, address string, page, pageSize int) ([]*neo4jstore.Node, int, error) {
	skip, limit := neo4jstore.BuildLimitOffset(page, pageSize)
	result, err := s.session.Run(ctx, neo4jstore.FindMentionsByAddressQuery, map[string]interface{}{
		"address": address,
		"skip":    skip,
		"limit":   limit,
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := neo4jstore.Count(ctx, s.session, neo4jstore.FindMentionsByAddressCountQuery, map[string]interface{}{
		"address": address,
	}, "t")
	if err != nil {
		return nil, 0, err
	}

	items, err := neo4jstore.GetItems(ctx, result, "m")
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (s *GenericStoreImpl) FindContactByAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindContactByAddressQuery, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "c")
}

func (s *GenericStoreImpl) findReportedRiskScore(ctx context.Context, address string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindAddressByHashQuery, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "a")
}

// FindWalletForAddress Сущность кошелек для адреса
func (s *GenericStoreImpl) FindWalletForAddress(ctx context.Context, address string) (*neo4jstore.Node, error) {
	result, err := s.session.Run(ctx, neo4jstore.FindWalletForAddressQuery, map[string]interface{}{
		"address": address,
	})
	if err != nil {
		return nil, err
	}

	return neo4jstore.GetItem(ctx, result, "w")
}

// Risk reported и calculated риски
func (s *GenericStoreImpl) Risk(ctx context.Context, address string) (*neo4jstore.Risk, error) {
	node, err := s.findReportedRiskScore(ctx, address)
	if err != nil {
		return nil, err
	}

	r := &neo4jstore.Risk{}

	if node != nil {
		calculated, err := neo4jstore.CalculateRiskFromNode(node)
		if err != nil {
			return nil, err
		}

		if calculated != nil {
			r.Risk = &calculated.Risk
			r.Calculated = calculated
		}

		wallet, err := s.FindWalletForAddress(ctx, address)
		if err != nil {
			return nil, err
		}

		if wallet != nil {
			if category, ok := wallet.Props["category"]; ok {
				categoryNumber := cast.ToInt(category)

				cat, err := s.category.CategoryFindByNumber(ctx, categoryNumber)
				if err != nil {
					return nil, err
				}

				if cat == nil {
					return nil, errors.New("category not found")
				}

				fr := float64(cat.Risk)
				r.Risk = &fr
				r.Wallet = &neo4jstore.RiskData{
					Category: cat.Number,
					Risk:     float64(cat.Risk),
				}
			}
		}

		if category, ok := node.Props["category"]; ok {
			cat, err := s.category.CategoryFindByName(ctx, category.(string))
			if err != nil {
				return nil, err
			}

			if cat != nil {
				fr := float64(cat.Risk)
				r.Risk = &fr
				r.Reported = &neo4jstore.RiskData{
					Category: cat.Number,
					Risk:     float64(cat.Risk),
				}
			}
		}

		if categoryID, ok := node.Props["category_id"]; ok {
			cat, err := s.category.CategoryFindByNumber(ctx, cast.ToInt(categoryID))
			if err != nil {
				return nil, err
			}

			if cat != nil {
				fr := float64(cat.Risk)
				r.Risk = &fr
				r.Reported = &neo4jstore.RiskData{
					Category: cat.Number,
					Risk:     float64(cat.Risk),
				}
			}
		}
	}

	if r.Wallet == nil && r.Reported == nil && r.Calculated == nil {
		return nil, nil
	}

	return r, nil
}
