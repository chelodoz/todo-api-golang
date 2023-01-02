package errors

import (
	"fmt"
	"net/http"
)

//Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorTypes
const (
	Authorization        Type = "AUTHORIZATION"          // Authentication Failures -
	BadRequest           Type = "BAD_REQUEST"            // Validation errors / BadInput
	Conflict             Type = "CONFLICT"               // Already exists (eg, create account with existent email) - 409
	Internal             Type = "INTERNAL"               // Server (500) and fallback errors
	NotFound             Type = "NOT_FOUND"              // For not finding resource
	UnprocessableEntity  Type = "UNPROCESSABLE_ENTITY"   // Not able to decode the JSON request - 422
	PayloadTooLarge      Type = "PAYLOAD_TOO_LARGE"      // For uploading tons of JSON, or an image over the limit - 413
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"    // For long running handlers
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE" // for http 415
)

type ErrorResponse struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

/*
* Error "Factories"
 */

// NewAuthorization to create a 401
func NewAuthorization(reason string) *ErrorResponse {
	return &ErrorResponse{
		Type:    Authorization,
		Message: reason,
		Code:    http.StatusUnauthorized,
	}
}

// NewBadRequest to create 400 errors (validation, for example)
func NewBadRequest(reason string) *ErrorResponse {
	return &ErrorResponse{
		Type:    BadRequest,
		Message: reason,
		Code:    http.StatusBadRequest,
	}
}

// NewConflict to create an error for 409
func NewConflict(name string, value string) *ErrorResponse {
	return &ErrorResponse{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
		Code:    http.StatusConflict,
	}
}

// NewUnprocessableEntity to create an error for 422
func NewUnprocessableEntity() *ErrorResponse {
	return &ErrorResponse{
		Type:    UnprocessableEntity,
		Message: fmt.Sprintf("Unable to process the request"),
		Code:    http.StatusUnprocessableEntity,
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal() *ErrorResponse {
	return &ErrorResponse{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *ErrorResponse {
	return &ErrorResponse{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
		Code:    http.StatusNotFound,
	}
}

// NewPayloadTooLarge to create an error for 413
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *ErrorResponse {
	return &ErrorResponse{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
		Code:    http.StatusRequestEntityTooLarge,
	}
}

// NewServiceUnavailable to create an error for 503
func NewServiceUnavailable() *ErrorResponse {
	return &ErrorResponse{
		Type:    ServiceUnavailable,
		Message: fmt.Sprintf("Service unavailable or timed out"),
		Code:    http.StatusServiceUnavailable,
	}
}

// NewUnsupportedMediaType to create an error for 415
func NewUnsupportedMediaType(reason string) *ErrorResponse {
	return &ErrorResponse{
		Type:    UnsupportedMediaType,
		Message: reason,
		Code:    http.StatusUnsupportedMediaType,
	}
}
