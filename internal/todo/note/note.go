package note

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID          uuid.UUID `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
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
