package note

import (
	"errors"
	"log"
	"net/http"
	"todo-api-golang/pkg/error"
	"todo-api-golang/pkg/util"

	"github.com/google/uuid"

	"github.com/go-playground/validator"
)

type handler struct {
	service  Service
	validate *validator.Validate
}

func NewHandler(service Service) Handler {
	validator := validator.New()
	if err := validator.RegisterValidation("enum", ValidateEnum); err != nil {
		log.Printf("Failed registering handler validators: %v", err)
	}

	return &handler{
		service:  service,
		validate: validator,
	}
}

// swagger:route POST /notes Notes CreateNoteRequestWrapper
// Creates a new note
//
// Create a new note in a database
//
// responses:
// 201: CreateNoteResponse
// 400: ValidationErrorResponseWrapper
// 422: ErrorResponseWrapper
// 500: ErrorResponseWrapper

// Create handles POST requests and create a note into the data store
func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	var createNoteRequest CreateNoteRequest

	if err := util.ReadRequestBody(r, &createNoteRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := h.validate.Struct(&createNoteRequest); err != nil {
		util.WriteError(rw, error.NewValidationBadRequest(err.(validator.ValidationErrors)))
		return
	}

	newNote := &Note{
		Name:        createNoteRequest.Name,
		Description: createNoteRequest.Description,
	}
	note, err := h.service.Create(newNote, r.Context())

	if err != nil {
		util.WriteError(rw, error.NewInternal())
		return
	}

	noteResponse := &CreateNoteResponse{
		ID:          note.ID.String(),
		Name:        note.Name,
		Description: note.Description,
		Status:      note.Status,
	}

	util.WriteResponse(rw, http.StatusCreated, noteResponse)
}

// swagger:route GET /notes Notes Notes
// Returns a list of notes
//
// Returns a list of notes from the database
// responses:
// 200: GetNotesResponse
// 500: ErrorResponseWrapper

// GetAll handles GET requests and returns all the notes from the data store
func (h *handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetAll(r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrFoundingNote):
			util.WriteResponse(rw, http.StatusOK, &GetNotesResponse{})
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	var notesResponse GetNotesResponse

	for _, note := range notes {
		noteResponse := GetNoteResponse{
			ID:          note.ID.String(),
			Name:        note.Name,
			Description: note.Description,
			Status:      note.Status,
		}
		notesResponse = append(notesResponse, noteResponse)
	}

	util.WriteResponse(rw, http.StatusOK, &notesResponse)
}

// swagger:route GET /notes/{noteId} Notes NoteIdQueryParamWrapper
// Returns a single note
//
// Returns a single note from the database
// responses:
// 200: GetNoteResponse
// 500: ErrorResponseWrapper

// GetNote handles GET/{noteId} requests and returns a note from the data store
func (h *handler) GetById(rw http.ResponseWriter, r *http.Request) {
	noteId, err := util.GetUriParam(r, "noteId")
	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidNoteId.Error()))
		return
	}

	uid, err := uuid.Parse(noteId)
	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidNoteId.Error()))
		return
	}

	note, err := h.service.GetById(uid, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrFoundingNote):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	noteResponse := &GetNoteResponse{
		ID:          note.ID.String(),
		Name:        note.Name,
		Description: note.Description,
		Status:      note.Status,
	}

	util.WriteResponse(rw, http.StatusOK, noteResponse)
}

// swagger:route PATCH /notes/{noteId} Notes UpdateNoteRequestWrapper
// Update an existing note
//
// Update a new note in a database
//
// responses:
// 204: NoContentResponseWrapper
// 400: ValidationErrorResponseWrapper
// 422: ErrorResponseWrapper
// 500: ErrorResponseWrapper

// Update handles PATCH requests and updates a note into the data store
func (h *handler) Update(rw http.ResponseWriter, r *http.Request) {
	var updateNoteRequest UpdateNoteRequest
	noteId, err := util.GetUriParam(r, "noteId")
	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidNoteId.Error()))
		return
	}

	uid, err := uuid.Parse(noteId)
	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidNoteId.Error()))
		return
	}

	if err := util.ReadRequestBody(r, &updateNoteRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := h.validate.Struct(&updateNoteRequest); err != nil {
		util.WriteError(rw, error.NewValidationBadRequest(err.(validator.ValidationErrors)))
		return
	}

	updatedNote := &Note{
		ID:          uid,
		Name:        updateNoteRequest.Name,
		Description: updateNoteRequest.Description,
		Status:      updateNoteRequest.Status,
	}

	_, err = h.service.Update(updatedNote, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrFoundingNote):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
