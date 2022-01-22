package graph

import "service-users/graph/model"

var userRepository = []*model.User{
	{
		ID:   1,
		Name: "Marcelo",
	},
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
