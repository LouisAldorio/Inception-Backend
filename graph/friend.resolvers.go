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

func (r *friendResolver) User(ctx context.Context, obj *model.Friend) (*model.User, error) {
	return service.GetUserByUsername(obj.Username)
}

func (r *friendOpsResolver) Add(ctx context.Context, obj *model.FriendOps, friends []*string) (*model.User, error) {
	user := middleware.AuthContext(ctx)
	return service.AddFriend(friends, user.Username), nil
}

// Friend returns generated.FriendResolver implementation.
func (r *Resolver) Friend() generated.FriendResolver { return &friendResolver{r} }

// FriendOps returns generated.FriendOpsResolver implementation.
func (r *Resolver) FriendOps() generated.FriendOpsResolver { return &friendOpsResolver{r} }

type friendResolver struct{ *Resolver }
type friendOpsResolver struct{ *Resolver }
