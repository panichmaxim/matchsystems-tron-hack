package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/internal/tools"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *queryResolver) TronFindAddressByHash(ctx context.Context, address string) (*model.FindAddressByHashNodeResponse, error) {
	node, err := r.svc.TronFindAddressByHash(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.FindAddressByHashNodeResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.FindAddressByHashNodeResponse{Node: node}, nil
}

func (r *queryResolver) TronFindTransactionsByAddress(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.TronFindTransactionsByAddress(ctx, hash, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) TronFindTransactionByHash(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.TronFindTransactionByHash(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) TronFindIncomingTransactionAddress(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.TronFindIncomingTransactions(ctx, hash, page, pageSize)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) TronFindOutcomingTransactionAddress(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.TronFindOutcomingTransactions(ctx, hash, page, pageSize)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeListResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) TronFindBlockByTransaction(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.TronFindBlockByTransaction(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) TronFindBlockByHeight(ctx context.Context, height int) (*model.NodeEntityResponse, error) {
	node, err := r.svc.TronFindBlockByHeight(ctx, height)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) TronFindTransactionsInBlock(ctx context.Context, height int, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.TronFindTransactionsInBlock(ctx, height, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) TronFindBlockByHash(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.TronFindBlockByHash(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) TronFindMentionsByAddress(ctx context.Context, address string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.TronFindMentionsForAddress(ctx, address, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) TronFindContactByAddress(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.TronFindContactByAddress(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) TronRisk(ctx context.Context, address string) (*model.RiskResponse, error) {
	risk, err := r.svc.TronRisk(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.RiskResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.RiskResponse{Risk: risk}, nil
}

func (r *queryResolver) TronSearch(ctx context.Context, query string, page int, limit int, wildcard *bool) (*model.SearchResponse, error) {
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
