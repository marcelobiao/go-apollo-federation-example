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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	fmt.Println("mutationResolver-CreateUser")
	return service.CreateUser(input)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	fmt.Println("queryResolver-User")
	return service.GetUser(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println("queryResolver-Users")
	return service.GetUsers()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
