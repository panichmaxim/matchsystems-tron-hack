package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/rubin-dev/api/internal/graph/generated"
	"gitlab.com/rubin-dev/api/pkg/models"
	"gitlab.com/rubin-dev/api/pkg/neo4jstore"
)

func (r *calculatedRiskResolver) Items(ctx context.Context, obj *neo4jstore.CalculatedRisk) ([]*neo4jstore.CalculateItem, error) {
	var items []*neo4jstore.CalculateItem
	for _, item := range obj.Items {
		items = append(items, item)
	}

	return items, nil
}

func (r *nodeResolver) Props(ctx context.Context, obj *neo4jstore.Node) (interface{}, error) {
	return obj.Props, nil
}

func (r *riskDataResolver) Category(ctx context.Context, obj *neo4jstore.RiskData) (*models.Category, error) {
	return r.svc.CategoryFindByNumber(ctx, obj.Category)
}

// CalculatedRisk returns generated.CalculatedRiskResolver implementation.
func (r *Resolver) CalculatedRisk() generated.CalculatedRiskResolver {
	return &calculatedRiskResolver{r}
}

// Node returns generated.NodeResolver implementation.
func (r *Resolver) Node() generated.NodeResolver { return &nodeResolver{r} }

// RiskData returns generated.RiskDataResolver implementation.
func (r *Resolver) RiskData() generated.RiskDataResolver { return &riskDataResolver{r} }

type calculatedRiskResolver struct{ *Resolver }
type nodeResolver struct{ *Resolver }
type riskDataResolver struct{ *Resolver }
