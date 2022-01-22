package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"service-users/graph/generated"
	"service-users/graph/model"
	"service-users/graph/service"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	fmt.Println("entityResolver-FindUserByID")
	return service.GetUser(id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
