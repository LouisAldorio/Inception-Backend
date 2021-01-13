package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"

	"github.com/LouisAldorio/Testing-early-injection-directive/middleware"
)

func (r *commodityOpsResolver) Create(ctx context.Context, obj *model.CommodityOps, input *model.NewComodity) (*model.Comodity, error) {
	userClaim := middleware.AuthContext(ctx)
	return service.CommodityCreate(ctx, input, userClaim.Username), nil
}

func (r *commodityOpsResolver) Update(ctx context.Context, obj *model.CommodityOps, input *model.NewComodity) (*model.Comodity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *comodityResolver) User(ctx context.Context, obj *model.Comodity) (*model.User, error) {
	return service.GetUserByUsername(obj.Username)
}

func (r *comodityPaginationResolver) TotalItem(ctx context.Context, obj *model.ComodityPagination) (int, error) {
	return service.GetTotalCommodity(ctx)
}

func (r *comodityPaginationResolver) Nodes(ctx context.Context, obj *model.ComodityPagination) ([]*model.Comodity, error) {
	return service.ComodityGetList(ctx, obj.Limit, obj.Page)
}

// CommodityOps returns generated.CommodityOpsResolver implementation.
func (r *Resolver) CommodityOps() generated.CommodityOpsResolver { return &commodityOpsResolver{r} }

// Comodity returns generated.ComodityResolver implementation.
func (r *Resolver) Comodity() generated.ComodityResolver { return &comodityResolver{r} }

// ComodityPagination returns generated.ComodityPaginationResolver implementation.
func (r *Resolver) ComodityPagination() generated.ComodityPaginationResolver {
	return &comodityPaginationResolver{r}
}

type commodityOpsResolver struct{ *Resolver }
type comodityResolver struct{ *Resolver }
type comodityPaginationResolver struct{ *Resolver }
