package service

import (
	"sample-golang-api/entity"
	"sample-golang-api/repository"
)

type TodoService interface {
	CreateTodo(todo entity.Todo) (*entity.Todo, error)
	GetTodoById(id uint) (*entity.Todo, error)
	GetTodos() ([]*entity.Todo, error)
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: repository,
	}
}

func (service *todoService) CreateTodo(todo entity.Todo) (*entity.Todo, error) {

	return &entity.Todo{}, nil
}

func (service *todoService) GetTodoById(id uint) (*entity.Todo, error) {

	return &entity.Todo{}, nil
}
func (service *todoService) GetTodos() ([]*entity.Todo, error) {

	return []*entity.Todo{}, nil
}
