package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"test/graph/generated"
	"test/graph/model"
)

func (r *mutationResolver) CreateCar(ctx context.Context, input *model.CreateCarInput) (*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
