package graph

import "service-task/graph/model"

var taskRepository = []*model.Task{
	{
		ID:   1,
		Task: "Create POC",
	},
}

func GetTasks() ([]*model.Task, error) {
	return taskRepository, nil
}

func CreateTask(input model.NewTask) (*model.Task, error) {
	user := model.Task{
		ID:   len(taskRepository) + 1,
		Task: input.Task,
	}
	taskRepository = append(taskRepository, &user)

	return &user, nil
}
