package graph

import "service-task/graph/model"

var taskRepository = []*model.Task{
	{
		ID:   1,
		Task: "Create POC",
		User: &model.User{
			ID: 1,
		},
	},
}

func GetTasks() ([]*model.Task, error) {
	return taskRepository, nil
}

func CreateTask(input model.NewTask) (*model.Task, error) {
	user := model.Task{
		ID:   len(taskRepository) + 1,
		Task: input.Task,
		User: &model.User{
			ID: input.UserID,
		},
	}
	taskRepository = append(taskRepository, &user)

	return &user, nil
}
