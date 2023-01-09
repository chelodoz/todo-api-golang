package note

import (
	"errors"
	"net/http"
	"todo-api-golang/pkg/error"
	"todo-api-golang/pkg/util"

	"github.com/google/uuid"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	validate = validator.New()
	return &handler{
		service: service,
	}
}

// swagger:route POST /notes Notes createNoteRequestWrapper
// Creates a new note
//
// Create a new note in a database
//
// responses:
// 201: CreateNoteResponse
// 422: errorResponseWrapper

// Create handles POST requests and create a note into the data store
func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	var createNoteRequest CreateNoteRequest

	if err := util.ReadRequestBody(r, &createNoteRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&createNoteRequest); err != nil {
		util.WriteError(rw, error.NewBadRequest(err.Error()))
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
	}

	util.WriteResponse(rw, http.StatusCreated, noteResponse)
}

// swagger:route GET /notes Notes Notes
// Returns a list of notes
//
// Returns a list of notes from the database
// responses:
// 200: GetNotesResponse

// GetAll handles GET requests and returns all the notes from the data store
func (h *handler) GetAll(rw http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetAll(r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrNoteNotFound):
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
		}
		notesResponse = append(notesResponse, noteResponse)
	}

	util.WriteResponse(rw, http.StatusOK, &notesResponse)
}

// swagger:route GET /notes/{noteId} Notes noteIdQueryParamWrapper
// Returns a single note
//
// Returns a single note from the database
// responses:
// 200: GetNoteByIdResponse

// GetNote handles GET/{noteId} requests and returns a note from the data store
func (h *handler) GetById(rw http.ResponseWriter, r *http.Request) {
	noteId := util.GetUriParam(r, "noteId")
	uid, err := uuid.Parse(noteId)

	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidId.Error()))
		return
	}

	note, err := h.service.GetById(uid, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrNoteNotFound):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	noteResponse := &GetNoteByIdResponse{
		ID:          note.ID.String(),
		Name:        note.Name,
		Description: note.Description,
	}

	util.WriteResponse(rw, http.StatusOK, noteResponse)
}

// swagger:route PATCH /notes/{noteId} Notes updateNoteRequestWrapper
// Update an existing note
//
// Update a new note in a database
//
// responses:
// 204: noContentResponseWrapper
// 422: errorResponseWrapper

// Update handles PATCH requests and updates a note into the data store
func (h *handler) Update(rw http.ResponseWriter, r *http.Request) {
	var updateNoteRequest UpdateNoteRequest
	noteId := util.GetUriParam(r, "noteId")
	uid, err := uuid.Parse(noteId)

	if err != nil {
		util.WriteError(rw, error.NewBadRequest(ErrInvalidId.Error()))
		return
	}
	if err := util.ReadRequestBody(r, &updateNoteRequest); err != nil {
		util.WriteError(rw, error.NewUnprocessableEntity())
		return
	}

	if err := validate.Struct(&updateNoteRequest); err != nil {
		util.WriteError(rw, error.NewBadRequest(err.Error()))
		return
	}

	updatedNote := &Note{
		ID:          uid,
		Name:        updateNoteRequest.Name,
		Description: updateNoteRequest.Description,
	}

	_, err = h.service.Update(updatedNote, r.Context())

	if err != nil {
		switch {
		case errors.Is(err, ErrNoteNotFound):
			util.WriteError(rw, error.NewNotFound())
		default:
			util.WriteError(rw, error.NewInternal())
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
