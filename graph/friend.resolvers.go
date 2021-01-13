package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *friendResolver) User(ctx context.Context, obj *model.Friend) (*model.User, error) {
	return service.GetUserByUsername(obj.Username)
}

// Friend returns generated.FriendResolver implementation.
func (r *Resolver) Friend() generated.FriendResolver { return &friendResolver{r} }

type friendResolver struct{ *Resolver }
