package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *scheduleResolver) InvolvedUsers(ctx context.Context, obj *model.Schedule) ([]*model.User, error) {
	return service.GetUserByWhereIn(obj.InvolvedUsersUsername), nil
}

func (r *scheduleOpsResolver) Create(ctx context.Context, obj *model.ScheduleOps, input model.NewSchedule) (*model.Schedule, error) {
	return service.CreateSchedule(ctx, input), nil
}

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

// ScheduleOps returns generated.ScheduleOpsResolver implementation.
func (r *Resolver) ScheduleOps() generated.ScheduleOpsResolver { return &scheduleOpsResolver{r} }

type scheduleResolver struct{ *Resolver }
type scheduleOpsResolver struct{ *Resolver }
