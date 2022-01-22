package service

import (
	"errors"
	"service-task/graph/model"
)

var taskRepository = []*model.Task{
	{
		ID:   1,
		Task: "Create POC",
		User: &model.User{
			ID: 1,
		},
	},
}

func GetTask(id int) (*model.Task, error) {
	for _, item := range taskRepository {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, errors.New("Task not found!")
}

func GetTasks() ([]*model.Task, error) {
	return taskRepository, nil
}

func GetUserTasks(id int) (*model.User, error) {
	tasks := []*model.Task{}
	for _, item := range taskRepository {
		if item.User.ID == id {
			tasks = append(tasks, item)
		}
	}

	return &model.User{Tasks: tasks}, nil
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
