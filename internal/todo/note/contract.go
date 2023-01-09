package note

// Create a note request
type CreateNoteRequest struct {
	// required: true
	Name string `json:"name" validate:"required"`
	// required: true
	Description string `json:"description" validate:"required"`
}

// Update a note request
type UpdateNoteRequest struct {
	// required: true
	Name string `json:"name" validate:"required"`
	// required: true
	Description string `json:"description" validate:"required"`
}

// Create a note response
// swagger:response CreateNoteResponse
type CreateNoteResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetNoteResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// List notes returns
// swagger:response GetNotesResponse
type GetNotesResponse []GetNoteResponse

// A single note returns in the response
// swagger:response GetNoteByIdResponse
type GetNoteByIdResponse CreateNoteResponse

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
