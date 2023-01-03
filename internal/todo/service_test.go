package todo

import (
	"testing"
	"todo-api-golang/internal/entity"
	"todo-api-golang/internal/mocks"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo_ValidTestInput_ShouldReturnCreatedTodoWithoutError(t *testing.T) {
	todo := &entity.Todo{
		Name:        "todo",
		Description: "todo description",
	}

	todoRepositoryMock := mocks.NewTodoRepository(t)
	todoService := NewTodoService(todoRepositoryMock)

	todoRepositoryMock.On("CreateTodo", mock.Anything, mock.Anything).Return(todo, nil)

	newTodo, err := todoService.CreateTodo(todo, nil)

	assert.Nil(t, err)
	assert.NotNil(t, newTodo)
}
