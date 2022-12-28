package repository

import (
	"sample-golang-api/entity"
)

type TodoRepository interface {
	CreateTodo(todo entity.Todo) (*entity.Todo, error)
	GetTodoById(id uint) (*entity.Todo, error)
	GetTodos() ([]entity.Todo, error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (todoRepository *todoRepository) CreateTodo(todo entity.Todo) (*entity.Todo, error) {
	return &entity.Todo{}, nil
}

func (todoRepository *todoRepository) GetTodoById(id uint) (*entity.Todo, error) {

	return &entity.Todo{}, nil
}

func (todoRepository *todoRepository) GetTodos() ([]entity.Todo, error) {
	var todos []entity.Todo
	return todos, nil
}
