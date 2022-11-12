package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *queryResolver) BtcFindContactByAddress(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindContactByAddress(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindTransactionByHash(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindTransactionByHash(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindAddressByHash(ctx context.Context, address string) (*model.FindAddressByHashNodeResponse, error) {
	node, err := r.svc.BtcFindAddressByHash(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.FindAddressByHashNodeResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.FindAddressByHashNodeResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindWalletForAddress(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindWalletForAddress(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcRisk(ctx context.Context, address string) (*model.RiskResponse, error) {
	risk, err := r.svc.BtcRisk(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RiskResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RiskResponse{Risk: risk}, nil
}

func (r *queryResolver) BtcFindBlockByNumber(ctx context.Context, height int) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindBlockByHeight(ctx, height)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindBlockByHash(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindBlockByHash(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindBlockByTransaction(ctx context.Context, txid string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindBlockByTransaction(ctx, txid)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindTransactionsByAddress(ctx context.Context, address string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindTransactionsByAddress(ctx, address, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindMentionsForAddress(ctx context.Context, address string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindMentionsForAddress(ctx, address, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindIncomingTransactions(ctx context.Context, txid string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindIncomingTransactions(ctx, txid, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindOutcomingTransactions(ctx context.Context, txid string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindOutcomingTransactions(ctx, txid, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindTransactionsInBlock(ctx context.Context, height int, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindTransactionsInBlock(ctx, height, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindTransactionsInBlockByHash(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindTransactionsInBlockByHash(ctx, hash, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcFindWalletByWid(ctx context.Context, wid string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.BtcFindWalletByWid(ctx, wid)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) BtcFindWalletAddresses(ctx context.Context, wid string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.BtcFindWalletAddresses(ctx, wid, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) BtcSearch(ctx context.Context, query string, page int, limit int, wildcard *bool) (*model.SearchResponse, error) {
	if wildcard == nil {
		wildcard = tools.Ptr[bool](true)
	}
	items, total, err := r.svc.Search(ctx, query, page, limit, *wildcard)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.SearchResponse{Errors: errs}, nil
		}
		return nil, err
	}

	return &model.SearchResponse{Edge: items, Total: &total}, nil
}
