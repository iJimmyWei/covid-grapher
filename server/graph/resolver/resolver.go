package resolver

import (
	"context"

	"github.com/ijimmywei/covid-grapher/db"
	"github.com/ijimmywei/covid-grapher/graph/generated"
	"github.com/ijimmywei/covid-grapher/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB db.DB
}

func (r *queryResolver) Records(ctx context.Context, skill string) ([]*model.Record, error) {
	return r.DB.GetRecords(skill)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
