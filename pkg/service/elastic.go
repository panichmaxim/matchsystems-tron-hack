package service

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/elastic"
)

type ElasticService interface {
	elastic.Client
}

var _ ElasticService = (*serviceImpl)(nil)

func (s *serviceImpl) Search(ctx context.Context, value string, page, limit int, wildcard bool) ([]*elastic.Entity, int, error) {
	if err := s.v.Search(ctx, value); err != nil {
		return nil, 0, err
	}

	return s.elk.Search(ctx, value, page, limit, wildcard)
}

func (s *serviceImpl) SearchCount(ctx context.Context, value string, wildcard bool) (int, error) {
	if err := s.v.SearchCount(ctx, value); err != nil {
		return -1, err
	}

	return s.elk.SearchCount(ctx, value, wildcard)
}
