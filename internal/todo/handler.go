package todo

import (
	"errors"
	"net/http"
	"todo-api-golang/internal/entity"
	appError "todo-api-golang/internal/error"
	"todo-api-golang/pkg/util"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type TodoHandler interface {
	CreateTodo(rw http.ResponseWriter, r *http.Request)
	GetTodoById(rw http.ResponseWriter, r *http.Request)
	GetTodos(rw http.ResponseWriter, r *http.Request)
	UpdateTodoById(rw http.ResponseWriter, r *http.Request)
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
// 422: errorResponseWrapper

//	CreateTodo handles POST requests and create a todo into the data store
func (h *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var createTodoRequest CreateTodoRequest

	if err := util.ReadRequestBody(r, &createTodoRequest); err != nil {

		util.WriteError(rw, appError.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&createTodoRequest); err != nil {
		util.WriteError(rw, appError.NewBadRequest(err.Error()))
		return
	}

	newTodo := &entity.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
	}
	todo, err := h.service.CreateTodo(newTodo, r.Context())

	if err != nil {
		util.WriteError(rw, appError.NewInternal())
		return
	}

	todoResponse := &CreateTodoResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
	}

	util.WriteResponse(rw, http.StatusCreated, todoResponse)
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
		util.WriteError(rw, appError.NewInternal())
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

	util.WriteResponse(rw, http.StatusOK, &todosResponse)
}

// swagger:route GET /todos/{todoId} Todos todoIdQueryParamWrapper
// Returns a single todo
//
// Returns a single todo from the database
// responses:
// 200: GetTodoByIdResponse

//	GetTodo handles GET/{todoId} requests and returns a todo from the data store
func (h *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId, err := util.GetIntId(r, "todoId")
	if err != nil {
		util.WriteError(rw, appError.NewBadRequest(err.Error()))
		return
	}

	todo, err := h.service.GetTodoById(todoId, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, appError.ErrTodoNotFound):
			util.WriteError(rw, appError.NewNotFound())
		default:
			util.WriteError(rw, appError.NewInternal())
		}
		return
	}

	todoResponse := &GetTodoByIdResponse{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
	}

	util.WriteResponse(rw, http.StatusOK, todoResponse)
}

// swagger:route PATCH /todos/{todoId} Todos updateTodoRequestWrapper
// Update an existing todo
//
// Update a new todo in a database
//
// responses:
// 204: noContentResponseWrapper
// 422: errorResponseWrapper

//	UpdateTodoById handles PATCH requests and updates a todo into the data store
func (h *todoHandler) UpdateTodoById(rw http.ResponseWriter, r *http.Request) {
	var updateTodoRequest UpdateTodoRequest
	todoId, err := util.GetIntId(r, "todoId")
	if err != nil {
		util.WriteError(rw, appError.NewBadRequest(err.Error()))
		return
	}

	if err := util.ReadRequestBody(r, &updateTodoRequest); err != nil {

		util.WriteError(rw, appError.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&updateTodoRequest); err != nil {
		util.WriteError(rw, appError.NewBadRequest(err.Error()))
		return
	}

	updatedTodo := &entity.Todo{
		ID:          todoId,
		Name:        updateTodoRequest.Name,
		Description: updateTodoRequest.Description,
	}
	_, err = h.service.UpdateTodo(updatedTodo, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, appError.ErrTodoNotFound):
			util.WriteError(rw, appError.NewNotFound())
		default:
			util.WriteError(rw, appError.NewInternal())
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
