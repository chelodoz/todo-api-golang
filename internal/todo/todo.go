package todo

import (
	"context"
	"net/http"
	"todo-api-golang/internal/entity"

	"github.com/google/uuid"
)

type TodoHandler interface {
	CreateTodo(rw http.ResponseWriter, r *http.Request)
	GetTodoById(rw http.ResponseWriter, r *http.Request)
	GetTodos(rw http.ResponseWriter, r *http.Request)
	UpdateTodoById(rw http.ResponseWriter, r *http.Request)
}

type TodoRepository interface {
	CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error)
	GetTodoById(id uuid.UUID, ctx context.Context) (*entity.Todo, error)
	GetTodos(ctx context.Context) ([]entity.Todo, error)
	UpdateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error)
}

type TodoService interface {
	CreateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error)
	GetTodoById(id uuid.UUID, ctx context.Context) (*entity.Todo, error)
	GetTodos(ctx context.Context) ([]entity.Todo, error)
	UpdateTodo(todo *entity.Todo, ctx context.Context) (*entity.Todo, error)
}
