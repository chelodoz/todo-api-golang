package handler

import (
	"net/http"
	"sample-golang-api/service"

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
func (u *todoHandler) CreateTodo(rw http.ResponseWriter, r *http.Request) {

}

//  GetTodos handles GET requests and returns all the todos from the data store
func (u *todoHandler) GetTodos(rw http.ResponseWriter, r *http.Request) {

}

//	GetTodo GET/{todoId} GET requests and returns a todo from the data store
func (u *todoHandler) GetTodoById(rw http.ResponseWriter, r *http.Request) {

}
