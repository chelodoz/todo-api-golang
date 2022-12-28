package service

import (
	"sample-golang-api/contract"
	"sample-golang-api/repository"
)

type TodoService interface {
	CreateTodo(todo contract.CreateTodoRequest) (*contract.GetTodoByIdResponse, error)
	GetTodos() (*contract.GetTodosResponse, error)
	GetTodoById(ID uint) (*contract.GetTodoByIdResponse, error)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: repository,
	}
}

func (service *todoService) CreateTodo(createTodoRequest contract.CreateTodoRequest) (*contract.GetTodoByIdResponse, error) {

	return &contract.GetTodoByIdResponse{}, nil
}

func (service *todoService) GetTodoById(ID uint) (*contract.GetTodoByIdResponse, error) {

	return &contract.GetTodoByIdResponse{}, nil
}
func (service *todoService) GetTodos() (*contract.GetTodosResponse, error) {

	return &contract.GetTodosResponse{}, nil
}
