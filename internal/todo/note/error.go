package note

import (
	"errors"
)

var (
	ErrInvalidNoteId  = errors.New("error invalid id")
	ErrCreatingNoteId = errors.New("error creating note id")
	ErrDecodingNote   = errors.New("error decoding note")
	ErrUpdatingNote   = errors.New("error updating note")
	ErrCreatingNote   = errors.New("error creating note")
	ErrFoundingNote   = errors.New("error founding note")
)
