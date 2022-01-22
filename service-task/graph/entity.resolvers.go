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

func (r *entityResolver) FindTaskByID(ctx context.Context, id int) (*model.Task, error) {
	fmt.Println("entityResolver-FindTaskByID")
	return service.GetTask(id)
}

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	fmt.Println("entityResolver-FindUserByID")
	return service.GetUserTasks(id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
