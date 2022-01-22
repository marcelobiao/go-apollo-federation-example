package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"service-task/graph/generated"
	"service-task/graph/model"
	"service-task/graph/service"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	fmt.Println("mutation-CreateTask")
	return service.CreateTask(input)
}

func (r *queryResolver) Task(ctx context.Context, id int) (*model.Task, error) {
	fmt.Println("queryResolver-Task")
	return service.GetTask(id)
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	fmt.Println("queryResolver-Tasks")
	return service.GetTasks()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
