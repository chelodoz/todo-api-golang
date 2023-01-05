package entity

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
}
