package errors

import "fmt"

// Code represents a stable error code for API responses.
type Code string

const (
CodeInternal       Code = "INTERNAL_ERROR"
CodeNotFound       Code = "NOT_FOUND"
CodeInvalidPayload Code = "INVALID_PAYLOAD"
)

// APIError encapsulates the structured error payload returned to clients.
type APIError struct {
Code    Code        `json:"code"`
Message string      `json:"message"`
Details interface{} `json:"details,omitempty"`
}

// Error satisfies the error interface.
func (e APIError) Error() string {
return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// New creates a new APIError with the provided code and message.
func New(code Code, message string, details interface{}) APIError {
return APIError{Code: code, Message: message, Details: details}
}
