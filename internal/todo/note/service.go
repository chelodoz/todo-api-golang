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

func (s *service) Create(note *Note, ctx context.Context) (*Note, error) {
	newNote, err := s.repository.Create(note, ctx)

	if err != nil {
		return nil, err
	}

	return newNote, nil
}
func (s *service) Update(note *Note, ctx context.Context) (*Note, error) {
	newNote, err := s.repository.Update(note, ctx)

	if err != nil {
		return nil, err
	}

	return newNote, nil
}

func (s *service) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {

	note, err := s.repository.GetById(id, ctx)

	if err != nil {
		return nil, err
	}

	return note, nil
}
func (s *service) GetAll(ctx context.Context) ([]Note, error) {

	notes, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return notes, nil
}
