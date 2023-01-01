package todo

import (
	"fmt"
	"sample-golang-api/internal/entity"
	"time"
)

var ErrTodoNotFound = fmt.Errorf("todo not found")

type TodoRepository interface {
	CreateTodo(todo entity.Todo) (*entity.Todo, error)
	GetTodoById(id uint) (*entity.Todo, error)
	GetTodos() ([]*entity.Todo, error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (todoRepository *todoRepository) CreateTodo(todo entity.Todo) (*entity.Todo, error) {
	maxID := todos[len(todos)-1].ID
	todo.ID = maxID + 1
	todos = append(todos, &todo)
	return &todo, nil
}

func (todoRepository *todoRepository) GetTodoById(id uint) (*entity.Todo, error) {
	i := findIndexById(id)
	if i == -1 {
		return nil, ErrTodoNotFound
	}
	return todos[i], nil
}

func (todoRepository *todoRepository) GetTodos() ([]*entity.Todo, error) {
	return todos, nil
}

var todos = []*entity.Todo{
	{
		ID:          1,
		Name:        "Setup stand up",
		Description: "Create an invitation for the daily stand up",
		CreatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Coffee",
		Description: "Grab some coffee to the team",
		CreatedOn:   time.Now().UTC().String(),
	},
}

func findIndexById(id uint) int {
	for i, p := range todos {
		if uint(p.ID) == id {
			return i
		}
	}

	return -1
}
