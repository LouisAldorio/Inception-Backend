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

func (r *userResolver) Products(ctx context.Context, obj *model.User) ([]*model.Comodity, error) {
	return service.GetCommoditiesByUsername(ctx,obj.Username),nil
}

func (r *userOpsResolver) Register(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.LoginResponse, error) {
	return service.Register(ctx, input)
}

func (r *userOpsResolver) Login(ctx context.Context, obj *model.UserOps, input model.LoginUser) (*model.LoginResponse, error) {
	user := utils.ForContext(ctx)
	fmt.Println(user)
	return service.Login(ctx, input)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
