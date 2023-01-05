package apperror

import (
	"fmt"
)

var (
	ErrInvalidId    = fmt.Errorf("invalid id")
	ErrTodoNotFound = fmt.Errorf("todo not found")
)
