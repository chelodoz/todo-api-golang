package todo

import (
	"net/http"
	"sample-golang-api/internal/entity"
	"sample-golang-api/internal/error"
	"sample-golang-api/internal/platform"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type TodoHandler interface {
	CreateTodo(rw http.ResponseWriter, r *http.Request)
	GetTodoById(rw http.ResponseWriter, r *http.Request)
	GetTodos(rw http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	service TodoService
}

func NewTodoHandler(service TodoService) TodoHandler {
	validate = validator.New()
	return &todoHandler{
		service: service,
	}
}

//	CreateTodo handles POST requests and create a todo into the data store
func (h *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var createTodoRequest CreateTodoRequest

	if err := platform.ReadRequestBody(r, &createTodoRequest); err != nil {
		platform.WriteResponse(rw, http.StatusBadRequest, error.ErrorResponse{Message: err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(entity.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
	})

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, error.ErrorResponse{Message: err.Error()})
		return
	}
	platform.WriteResponse(rw, http.StatusCreated, CreateTodoResponse{
		Id:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
	})
}

//  GetTodos handles GET requests and returns all the todos from the data store
func (h *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos()

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, error.ErrorResponse{Message: err.Error()})
		return
	}

	platform.WriteResponse(rw, http.StatusOK, todos)
}

//	GetTodo GET/{todoId} GET requests and returns a todo from the data store
func (h *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId := platform.GetIntId(r, "todoId")
	todos, err := h.service.GetTodoById(todoId)

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, error.ErrorResponse{Message: err.Error()})
		return
	}

	platform.WriteResponse(rw, http.StatusOK, todos)
}
