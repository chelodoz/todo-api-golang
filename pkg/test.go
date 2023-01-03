package pkg

import (
	"todo-api-golang/internal/entity"
)

type TestService interface {
	CreateTodo(todo entity.Todo) (entity.Todo, error)
}
