package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"service-task/graph/generated"
	"service-task/graph/model"
)

func (r *entityResolver) FindTaskByID(ctx context.Context, id int) (*model.Task, error) {
	for _, item := range taskRepository {
		if item.ID == id {
			return item, nil
		}
	}

	return &model.Task{}, errors.New("Task not found!")
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
