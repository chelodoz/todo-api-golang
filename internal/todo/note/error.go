package note

import (
	"fmt"
)

var (
	ErrInvalidId    = fmt.Errorf("invalid id")
	ErrNoteNotFound = fmt.Errorf("todo not found")
)
