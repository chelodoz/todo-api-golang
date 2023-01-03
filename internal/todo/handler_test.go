package todo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api-golang/internal/entity"
	"todo-api-golang/internal/mocks"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/mock"

	"github.com/gorilla/mux"
)

type CreateTodoTestSuite struct {
	suite.Suite
	router          *mux.Router
	todoHandlerTest TodoHandler
	mockTodoService *mocks.TodoService
}

func TestCreateTodoTestSuite(t *testing.T) {
	suite.Run(t, &CreateTodoTestSuite{})
}

func (suite *CreateTodoTestSuite) SetupSuite() {
	suite.mockTodoService = mocks.NewTodoService(suite.T())
	suite.todoHandlerTest = NewTodoHandler(suite.mockTodoService)
	suite.router = mux.NewRouter()
	suite.router.HandleFunc("/todos", suite.todoHandlerTest.CreateTodo).Methods(http.MethodPost)
}

func (suite *CreateTodoTestSuite) TestCreateTodo_ValidTestInput_ShouldReturn201Created() {
	// Arrange
	var request = strings.NewReader(`{
		"name": "test",
		"description": "description"
	}`)

	suite.mockTodoService.On("CreateTodo", mock.Anything, mock.Anything).Return(&entity.Todo{
		ID:          3,
		Name:        "test",
		Description: "description",
		CreatedAt:   "",
	}, nil)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/todos", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *CreateTodoTestSuite) ExecuteRequest(req *http.Request, router *mux.Router) *http.Response {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	res := recorder.Result()
	return res
}
func (suite *CreateTodoTestSuite) TestCreateTodo_InvalidInputWithMissingDescription_ShouldReturn400BadRequest() {
	// Arrange
	var request = strings.NewReader(`{
		"name": "test"
	}`)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/todos", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusBadRequest, res.StatusCode)
}

func (suite *CreateTodoTestSuite) TestCreateTodo_InvalidInputWithWrongJsonFormat_ShouldReturn422UnprocessableEntity() {
	// Arrange
	var request = strings.NewReader(`{
		"name": ["test"]
	}`)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/todos", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
}
