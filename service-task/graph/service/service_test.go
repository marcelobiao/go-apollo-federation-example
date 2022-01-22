package service_test

import (
	"fmt"
	"service-task/graph/model"
	"service-task/graph/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTask(t *testing.T) {
	b := true
	input := model.UpdateTaskDto{
		ID:       1,
		Finished: &b,
	}

	resp, err := service.UpdateTask(input)

	assert.Equal(t, nil, err)

	fmt.Println(resp)

}
