// Package classification Todo API
//
// Documentation for Todo API
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
package todo

// swagger:parameters createTodoRequestWrapper
type createTodoRequestWrapper struct {
	// in: body
	// required: true
	Body CreateTodoRequest
}

// swagger:parameters todoIdQueryParamWrapper
type todoIdQueryParamWrapper struct {
	// The id of the todo for which the operation relates
	// in: path
	// required: true
	ID uint `json:"todoId"`
}

// swagger:response errorResponseWrapper
type errorResponseWrapper struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
