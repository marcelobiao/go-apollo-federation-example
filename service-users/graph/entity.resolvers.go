package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"service-users/graph/generated"
	"service-users/graph/model"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	for _, item := range userRepository {
		if item.ID == id {
			return item, nil
		}
	}

	return &model.User{}, errors.New("Task not found!")
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
