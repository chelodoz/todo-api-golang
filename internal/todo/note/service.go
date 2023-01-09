package note

import (
	"context"

	"github.com/google/uuid"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (service *service) Create(note *Note, ctx context.Context) (*Note, error) {
	newNote, err := service.repository.Create(note, ctx)

	if err != nil {
		return nil, err
	}

	return newNote, nil
}
func (service *service) Update(note *Note, ctx context.Context) (*Note, error) {
	newNote, err := service.repository.Update(note, ctx)

	if err != nil {
		return nil, err
	}

	return newNote, nil
}

func (service *service) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {

	noteById, err := service.repository.GetById(id, ctx)

	if err != nil {
		return nil, err
	}

	return noteById, nil
}
func (service *service) GetAll(ctx context.Context) ([]Note, error) {

	notes, err := service.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return notes, nil
}
