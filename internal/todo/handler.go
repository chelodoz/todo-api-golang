package todo

import (
	"errors"
	"log"
	"net/http"
	appError "todo-api-golang/internal/apperror"
	"todo-api-golang/internal/entity"
	"todo-api-golang/pkg/error"
	"todo-api-golang/pkg/util"

	"github.com/google/uuid"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

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

// CreateTodo handles POST requests and create a todo into the data store
func (h *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var createTodoRequest CreateTodoRequest

	log.Printf("request received")

	if err := util.ReadRequestBody(r, &createTodoRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&createTodoRequest); err != nil {
		util.WriteError(rw, error.NewBadRequest(err.Error()))
		return
	}

	newTodo := &entity.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
	}
	todo, err := h.service.CreateTodo(newTodo, r.Context())

	if err != nil {
		util.WriteError(rw, error.NewInternal())
		return
	}

	todoResponse := &CreateTodoResponse{
		ID:          todo.ID.String(),
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

// GetTodos handles GET requests and returns all the todos from the data store
func (h *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos(r.Context())

	if err != nil {
		switch {
		case errors.Is(err, appError.ErrTodoNotFound):
			util.WriteResponse(rw, http.StatusOK, &GetTodosResponse{})
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	var todosResponse GetTodosResponse

	for _, todo := range todos {
		todoResponse := GetTodoResponse{
			ID:          todo.ID.String(),
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

// GetTodo handles GET/{todoId} requests and returns a todo from the data store
func (h *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId := util.GetUriParam(r, "todoId")
	uid, err := uuid.Parse(todoId)

	if err != nil {
		util.WriteError(rw, error.NewBadRequest(appError.ErrInvalidId.Error()))
		return
	}

	todo, err := h.service.GetTodoById(uid, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, appError.ErrTodoNotFound):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	todoResponse := &GetTodoByIdResponse{
		ID:          todo.ID.String(),
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

// UpdateTodoById handles PATCH requests and updates a todo into the data store
func (h *todoHandler) UpdateTodoById(rw http.ResponseWriter, r *http.Request) {
	var updateTodoRequest UpdateTodoRequest
	todoId := util.GetUriParam(r, "todoId")
	uid, err := uuid.Parse(todoId)

	if err != nil {
		util.WriteError(rw, error.NewBadRequest(appError.ErrInvalidId.Error()))
		return
	}
	if err := util.ReadRequestBody(r, &updateTodoRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&updateTodoRequest); err != nil {
		util.WriteError(rw, error.NewBadRequest(err.Error()))
		return
	}

	updatedTodo := &entity.Todo{
		ID:          uid,
		Name:        updateTodoRequest.Name,
		Description: updateTodoRequest.Description,
	}

	_, err = h.service.UpdateTodo(updatedTodo, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, appError.ErrTodoNotFound):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
