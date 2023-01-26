package note

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID          uuid.UUID `bson:"_id,omitempty"`
	Name        string    `bson:"name,omitempty"`
	Description string    `bson:"description,omitempty"`
	Status      Status    `bson:"status,omitempty"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
}

// swagger:enum Status
type Status string

const (
	Todo       Status = "To Do"
	InProgress Status = "In Progress"
	Done       Status = "Done"
)

func (s Status) IsValid() bool {
	switch s {
	case Todo, InProgress, Done:
		return true
	default:
		return false
	}
}

type Handler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	GetById(rw http.ResponseWriter, r *http.Request)
	GetAll(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
}

type Repository interface {
	Create(note *Note, ctx context.Context) (*Note, error)
	GetById(id uuid.UUID, ctx context.Context) (*Note, error)
	GetAll(ctx context.Context) ([]Note, error)
	Update(note *Note, ctx context.Context) (*Note, error)
}

type Service interface {
	Create(note *Note, ctx context.Context) (*Note, error)
	GetById(id uuid.UUID, ctx context.Context) (*Note, error)
	GetAll(ctx context.Context) ([]Note, error)
	Update(note *Note, ctx context.Context) (*Note, error)
}
