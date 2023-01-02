package todo

import (
	"context"
	"todo-api-golang/internal/entity"
)

type TodoService interface {
	CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error)
	GetTodoById(id uint, ctx context.Context) (*entity.Todo, error)
	GetTodos(ctx context.Context) ([]entity.Todo, error)
}

type todoService struct {
	todoRepository TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return &todoService{
		todoRepository: repository,
	}
}

func (service *todoService) CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error) {
	newTodo, err := service.todoRepository.CreateTodo(todo, ctx)

	if err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (service *todoService) GetTodoById(id uint, ctx context.Context) (*entity.Todo, error) {

	todoById, err := service.todoRepository.GetTodoById(id, ctx)

	if err != nil {
		return nil, err
	}

	return todoById, nil
}
func (service *todoService) GetTodos(ctx context.Context) ([]entity.Todo, error) {

	todos, err := service.todoRepository.GetTodos(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
