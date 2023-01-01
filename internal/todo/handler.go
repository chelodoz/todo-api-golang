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

// swagger:route POST /todos Todos createTodoRequestWrapper
// Creates a new todo
//
// Create a new todo in a database
//
// responses:
// 201: CreateTodoResponse

//	CreateTodo handles POST requests and create a todo into the data store
func (h *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var createTodoRequest CreateTodoRequest

	if err := platform.ReadRequestBody(r, &createTodoRequest); err != nil {
		platform.WriteResponse(rw, http.StatusBadRequest, &error.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	newTodo := &entity.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
	}
	todo, err := h.service.CreateTodo(newTodo, r.Context())

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, &error.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	todoResponse := &CreateTodoResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
	}

	platform.WriteResponse(rw, http.StatusCreated, todoResponse)
}

// swagger:route GET /todos Todos Todos
// Returns a list of todos
//
// Returns a list of todos from the database
// responses:
// 200: GetTodosResponse

//  GetTodos handles GET requests and returns all the todos from the data store
func (h *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos(r.Context())

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, &error.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	var todosResponse GetTodosResponse

	for _, todo := range todos {
		todoResponse := CreateTodoResponse{
			ID:          todo.ID,
			Name:        todo.Name,
			Description: todo.Description,
		}
		todosResponse = append(todosResponse, todoResponse)
	}

	platform.WriteResponse(rw, http.StatusOK, &todosResponse)
}

// swagger:route GET /todos/{todoId} Todos todoIdQueryParamWrapper
// Returns a single todo
//
// Returns a single todo from the database
// responses:
// 200: GetTodoByIdResponse

//	GetTodo handles GET/{todoId} requests and returns a todo from the data store
func (h *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId := platform.GetIntId(r, "todoId")
	todo, err := h.service.GetTodoById(todoId, r.Context())

	if err != nil {
		platform.WriteResponse(rw, http.StatusInternalServerError, &error.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	todoResponse := &GetTodoByIdResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
	}

	platform.WriteResponse(rw, http.StatusOK, todoResponse)
}
