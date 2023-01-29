package error

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

// Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorTypes
const (
	Authorization        Type = "AUTHORIZATION"          // Authentication Failures -
	BadRequest           Type = "BAD_REQUEST"            // Validation errors / BadInput
	Conflict             Type = "CONFLICT"               // Resource already exists 409
	Internal             Type = "INTERNAL"               // Server (500) and fallback errors
	NotFound             Type = "NOT_FOUND"              // For not finding resource
	TooManyRequest       Type = "TOO_MANY_REQUEST"       // For not finding resource
	UnprocessableEntity  Type = "UNPROCESSABLE_ENTITY"   // Not able to decode the JSON request - 422
	PayloadTooLarge      Type = "PAYLOAD_TOO_LARGE"      // For uploading tons of JSON, or an image over the limit - 413
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"    // For long running handlers
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE" // For http 415
)

type ErrorResponse struct {
	Type    Type              `json:"type"`
	Message string            `json:"message"`
	Code    int               `json:"code"`
	Detail  string            `json:"detail,omitempty"`
	Errors  []validationError `json:"errors,omitempty"`
}

// NewAuthorization to create a 401
func NewAuthorization(reason string) *ErrorResponse {
	return &ErrorResponse{
		Type:    Authorization,
		Message: reason,
		Code:    http.StatusUnauthorized,
	}
}

// NewBadRequest to create 400 errors
func NewBadRequest(reason string) *ErrorResponse {
	return &ErrorResponse{
		Type:    BadRequest,
		Message: reason,
		Code:    http.StatusBadRequest,
	}
}

type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// NewValidationBadRequest to create 400 validation errors
func NewValidationBadRequest(ve validator.ValidationErrors) *ErrorResponse {
	var validationErrors []validationError
	for _, fe := range ve {
		validationErrors = append(validationErrors, validationError{fe.Field(), msgForTag(fe.Tag())})
	}

	return &ErrorResponse{
		Type:    BadRequest,
		Message: "One of the request inputs is not valid.",
		Code:    http.StatusBadRequest,
		Errors:  validationErrors,
	}
}

// msgForTag create a simple mapper from the validators tags to a custom message
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "enum":
		return "Invalid enum value."
	default:
		return tag
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
		Message: "Unable to process the request.",
		Code:    http.StatusUnprocessableEntity,
	}
}

// NewInternal for 500 errors and unknown errors
func NewInternal(detail string) *ErrorResponse {
	return &ErrorResponse{
		Type:    Internal,
		Message: "Internal server error.",
		Detail:  detail,
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFound to create an error for 404
func NewNotFound() *ErrorResponse {
	return &ErrorResponse{
		Type:    NotFound,
		Message: "The specified resource does not exist.",
		Code:    http.StatusNotFound,
	}
}

// NewPayloadTooLarge to create an error for 413
func NewPayloadTooLarge(maxBodySize int64, contentLength int64) *ErrorResponse {
	return &ErrorResponse{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v.", maxBodySize, contentLength),
		Code:    http.StatusRequestEntityTooLarge,
	}
}

// NewServiceUnavailable to create an error for 503
func NewServiceUnavailable() *ErrorResponse {
	return &ErrorResponse{
		Type:    ServiceUnavailable,
		Message: "Service unavailable or timed out.",
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
