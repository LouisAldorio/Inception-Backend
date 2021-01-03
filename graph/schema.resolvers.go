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

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	return utils.GetUserByUsername(username)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := utils.ForContext(ctx)
	fmt.Println(user.Username)

	return nil, nil
}
