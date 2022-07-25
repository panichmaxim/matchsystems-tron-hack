package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/model"
	"gitlab.com/rubin-dev/api/pkg/validator"
)

func (r *queryResolver) EthFindAddressByHash(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindAddressByHash(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindTransactionsByAddress(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.EthFindTransactionsByAddress(ctx, hash, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) EthFindTransactionByHash(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindTransactionByHash(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindIncomingTransactionAddress(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindIncomingTransactionAddress(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindOutcomingTransactionAddress(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindOutcomingTransactionAddress(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindBlockByTransaction(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindBlockByTransaction(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindBlockByHeight(ctx context.Context, height string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindBlockByHeight(ctx, height)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindTransactionsInBlock(ctx context.Context, height string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.EthFindTransactionsInBlock(ctx, height, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) EthFindAllInputAndOutputTransactions(ctx context.Context, hash string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.EthFindAllInputAndOutputTransactions(ctx, hash, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) EthFindBlockByHash(ctx context.Context, hash string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindBlockByHash(ctx, hash)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindMentionsByAddress(ctx context.Context, address string, page int, pageSize int) (*model.NodeListResponse, error) {
	nodes, total, err := r.svc.EthFindMentionsByAddress(ctx, address, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &model.NodeListResponse{
		Total: &total,
		Edge:  nodes,
	}, nil
}

func (r *queryResolver) EthFindContactByAddress(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindContactByAddress(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}

func (r *queryResolver) EthFindRiskScoreByAddress(ctx context.Context, address string) (*model.NodeEntityResponse, error) {
	node, err := r.svc.EthFindRiskScoreByAddress(ctx, address)
	if err != nil {
		if errs, ok := err.(validator.Errors); ok {
			return &model.NodeEntityResponse{Errors: errs}, nil
		}

		return nil, err
	}

	return &model.NodeEntityResponse{Node: node}, nil
}
