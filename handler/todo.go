package handler

import (
	"net/http"
	"sample-golang-api/contract"
	"sample-golang-api/entity"
	"sample-golang-api/service"
	"sample-golang-api/util"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type TodoHandler interface {
	CreateTodo(rw http.ResponseWriter, r *http.Request)
	GetTodoById(rw http.ResponseWriter, r *http.Request)
	GetTodos(rw http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	validate = validator.New()
	return &todoHandler{
		service: service,
	}
}

//	CreateTodo handles POST requests and create a todo into the data store
func (h *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	var createTodoRequest contract.CreateTodoRequest

	if err := util.ReadRequestBody(r, &createTodoRequest); err != nil {
		util.WriteResponse(rw, http.StatusBadRequest, contract.AppError{Message: err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(entity.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
	})

	if err != nil {
		util.WriteResponse(rw, http.StatusInternalServerError, contract.AppError{Message: err.Error()})
		return
	}
	util.WriteResponse(rw, http.StatusCreated, contract.CreateTodoResponse{
		Id:          todo.Id,
		Name:        todo.Name,
		Description: todo.Description,
	})
}

//  GetTodos handles GET requests and returns all the todos from the data store
func (h *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetTodos()

	if err != nil {
		util.WriteResponse(rw, http.StatusInternalServerError, contract.AppError{Message: err.Error()})
		return
	}

	util.WriteResponse(rw, http.StatusOK, todos)
}

//	GetTodo GET/{todoId} GET requests and returns a todo from the data store
func (h *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId := util.GetIntId(r, "todoId")
	todos, err := h.service.GetTodoById(todoId)

	if err != nil {
		util.WriteResponse(rw, http.StatusInternalServerError, contract.AppError{Message: err.Error()})
		return
	}

	util.WriteResponse(rw, http.StatusOK, todos)
}
