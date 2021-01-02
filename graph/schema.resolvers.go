package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
	"myapp/utils"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Commodity(ctx context.Context) (*model.CommodityOps, error) {
	return &model.CommodityOps{}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := utils.ForContext(ctx)
	fmt.Println(user.Username)

	return nil, nil
}

func (r *queryResolver) Comodities(ctx context.Context, limit *int, page *int) (*model.ComodityPagination, error) {
	return &model.ComodityPagination{
		Limit: limit,
		Page:  page,
	}, nil
}

func (r *queryResolver) UsersByRole(ctx context.Context, role string) ([]*model.User, error) {
	return service.GetUserByRole(role), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
