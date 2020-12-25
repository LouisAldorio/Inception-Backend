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

func (r *commodityOpsResolver) Create(ctx context.Context, obj *model.CommodityOps, input *model.NewComodity) (*model.Comodity, error) {
	user := utils.ForContext(ctx)
	return service.CommodityCreate(input, user), nil
}

func (r *comodityPaginationResolver) TotalItem(ctx context.Context, obj *model.ComodityPagination) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *comodityPaginationResolver) Nodes(ctx context.Context, obj *model.ComodityPagination) ([]*model.Comodity, error) {
	return service.ComodityGetList(ctx, obj.Limit, obj.Page)
}

// CommodityOps returns generated.CommodityOpsResolver implementation.
func (r *Resolver) CommodityOps() generated.CommodityOpsResolver { return &commodityOpsResolver{r} }

// ComodityPagination returns generated.ComodityPaginationResolver implementation.
func (r *Resolver) ComodityPagination() generated.ComodityPaginationResolver {
	return &comodityPaginationResolver{r}
}

type commodityOpsResolver struct{ *Resolver }
type comodityPaginationResolver struct{ *Resolver }
