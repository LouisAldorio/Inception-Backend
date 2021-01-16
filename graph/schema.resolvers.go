package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"

	"github.com/LouisAldorio/Testing-early-injection-directive/middleware"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Commodity(ctx context.Context) (*model.CommodityOps, error) {
	return &model.CommodityOps{}, nil
}

func (r *mutationResolver) Schedule(ctx context.Context) (*model.ScheduleOps, error) {
	return &model.ScheduleOps{}, nil
}

func (r *mutationResolver) Friends(ctx context.Context) (*model.FriendOps, error) {
	return &model.FriendOps{}, nil
}

func (r *queryResolver) UserByUsername(ctx context.Context, username string) (*model.User, error) {
	return service.GetUserByUsername(username)
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

func (r *queryResolver) ScheduleByUser(ctx context.Context) ([]*model.Schedule, error) {
	user := middleware.AuthContext(ctx)
	return service.GetSchedule(user.Username), nil
}

func (r *queryResolver) FriendList(ctx context.Context) ([]*model.Friend, error) {
	user := middleware.AuthContext(ctx)
	return service.GetFriendDetailByUsername(user.Username), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
