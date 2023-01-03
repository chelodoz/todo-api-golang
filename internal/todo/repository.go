package todo

import (
	"context"
	"time"
	"todo-api-golang/internal/entity"
	appError "todo-api-golang/internal/error"
)

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (todoRepository *todoRepository) CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error) {
	maxID := todos[len(todos)-1].ID
	todo.ID = maxID + 1
	todo.CreatedAt = time.Now().UTC().String()
	todos = append(todos, *todo)
	return todo, nil
}

func (todoRepository *todoRepository) GetTodoById(id uint, ctx context.Context) (*entity.Todo, error) {
	i := findIndexById(id)
	if i == -1 {
		return nil, appError.ErrTodoNotFound
	}
	return &todos[i], nil
}

func (todoRepository *todoRepository) GetTodos(ctx context.Context) ([]entity.Todo, error) {
	return todos, nil
}

var todos = []entity.Todo{
	{
		ID:          1,
		Name:        "Setup stand up",
		Description: "Create an invitation for the daily stand up",
		CreatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Coffee",
		Description: "Grab some coffee to the team",
		CreatedAt:   time.Now().UTC().String(),
	},
}

func (todoRepository *todoRepository) UpdateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error) {

	i := findIndexById(todo.ID)
	if i == -1 {
		return nil, appError.ErrTodoNotFound
	}

	todos[i] = *todo
	return todo, nil
}

func findIndexById(id uint) int {
	for i, p := range todos {
		if uint(p.ID) == id {
			return i
		}
	}
	return -1
}
