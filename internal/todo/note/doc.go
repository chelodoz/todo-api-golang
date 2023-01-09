// Package classification Todo API
//
// # Documentation for Todo API
//
// Schemes: http, https
// Host: localhost:8080
// BasePath: /api/v1
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package note

// swagger:parameters createNoteRequestWrapper
type createNoteRequestWrapper struct {
	// in: body
	// required: true
	Body CreateNoteRequest
}

// swagger:parameters updateNoteRequestWrapper
type updateNoteRequestWrapper struct {
	// The id of the note for which the operation relates
	// in: path
	// required: true
	ID string `json:"noteId"`
	// in: body
	// required: true
	Body UpdateNoteRequest
}

// swagger:parameters noteIdQueryParamWrapper
type noteIdQueryParamWrapper struct {
	// The id of the note for which the operation relates
	// in: path
	// required: true
	ID string `json:"noteId"`
}

// No content is returned by this API endpoint
// swagger:response noContentResponseWrapper
type noContentResponseWrapper struct {
}

// swagger:response errorResponseWrapper
type errorResponseWrapper struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
