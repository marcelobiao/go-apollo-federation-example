package graph

import (
	"errors"
	"service-users/graph/model"
)

var userRepository = []*model.User{
	{
		ID:   1,
		Name: "Marcelo",
	},
}

func GetUser(id int) (*model.User, error) {
	for _, item := range userRepository {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, errors.New("User not found!")
}

func GetUsers() ([]*model.User, error) {
	return userRepository, nil
}

func CreateUser(input model.NewUser) (*model.User, error) {
	user := model.User{
		ID:   len(userRepository) + 1,
		Name: input.Name,
	}
	userRepository = append(userRepository, &user)

	return &user, nil
}
