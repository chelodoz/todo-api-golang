package todo

import (
	"testing"
	"time"
	"todo-api-golang/internal/entity"
	"todo-api-golang/internal/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateTodoServiceTestSuite struct {
	suite.Suite
	todoService        TodoService
	mockTodoRepository *mocks.TodoRepository
}

func TestCreateTodoTestSuite(t *testing.T) {
	suite.Run(t, &CreateTodoHandlerTestSuite{})
}

func (suite *CreateTodoServiceTestSuite) SetupSuite() {
	suite.mockTodoRepository = mocks.NewTodoRepository(suite.T())
	suite.todoService = NewTodoService(suite.mockTodoRepository)
}
func (suite *CreateTodoServiceTestSuite) TestCreateTodo_ValidTestInput_ShouldReturnCreatedTodoWithoutError() {
	todo := &entity.Todo{
		ID:          uuid.New(),
		Name:        "todo",
		Description: "todo description",
		CreatedAt:   time.Now().UTC(),
	}

	suite.mockTodoRepository.On("CreateTodo", mock.Anything, mock.Anything).Return(todo, nil)

	newTodo, err := suite.todoService.CreateTodo(todo, nil)

	suite.Nil(err)
	suite.NotNil(newTodo)
}
