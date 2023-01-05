package todo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"todo-api-golang/internal/entity"
	"todo-api-golang/internal/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/mock"

	"github.com/gorilla/mux"
)

type CreateTodoHandlerTestSuite struct {
	suite.Suite
	router          *mux.Router
	todoHandlerTest TodoHandler
	mockTodoService *mocks.TodoService
}

func TestCreateTodoHandlerTestSuite(t *testing.T) {
	suite.Run(t, &CreateTodoHandlerTestSuite{})
}

func (suite *CreateTodoHandlerTestSuite) SetupSuite() {
	suite.mockTodoService = mocks.NewTodoService(suite.T())
	suite.todoHandlerTest = NewTodoHandler(suite.mockTodoService)
	suite.router = mux.NewRouter()
	suite.router.HandleFunc("/todos", suite.todoHandlerTest.CreateTodo).Methods(http.MethodPost)
}

func (suite *CreateTodoHandlerTestSuite) TestCreateTodo_ValidTestInput_ShouldReturn201Created() {
	// Arrange
	var request = strings.NewReader(`{
		"name": "test",
		"description": "description"
	}`)

	suite.mockTodoService.On("CreateTodo", mock.Anything, mock.Anything).Return(&entity.Todo{
		ID:          uuid.New(),
		Name:        "test",
		Description: "description",
		CreatedAt:   time.Now().UTC(),
	}, nil)

	// Act
	req, _ := http.NewRequest(http.MethodPost, "/todos", request)
	res := suite.ExecuteRequest(req, suite.router)

	// Assert

	suite.Equal(http.StatusCreated, res.StatusCode)
}

func (suite *CreateTodoHandlerTestSuite) ExecuteRequest(req *http.Request, router *mux.Router) *http.Response {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	res := recorder.Result()
	return res
}
func (suite *CreateTodoHandlerTestSuite) TestCreateTodo_InvalidInputWithMissingDescription_ShouldReturn400BadRequest() {
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

func (suite *CreateTodoHandlerTestSuite) TestCreateTodo_InvalidInputWithWrongJsonFormat_ShouldReturn422UnprocessableEntity() {
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
