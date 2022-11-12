package neo4jstore

import (
	"context"
	"gitlab.com/rubin-dev/api/pkg/models"
)

type CategoryRisk interface {
	CategoryFindByNumber(ctx context.Context, id int) (*models.Category, error)
	CategoryFindByName(ctx context.Context, name string) (*models.Category, error)
}
