package errors

import "net/http"

// ErrorCode ...
type ErrorCode int32

// OK ...
const (
	OK                  ErrorCode = http.StatusOK
	BadRequest          ErrorCode = http.StatusBadRequest
	Unauthorized        ErrorCode = http.StatusUnauthorized
	Forbidden           ErrorCode = http.StatusForbidden
	NotFound            ErrorCode = http.StatusNotFound
	MethodNotAllowed    ErrorCode = http.StatusMethodNotAllowed
	Timeout             ErrorCode = http.StatusRequestTimeout
	Conflict            ErrorCode = http.StatusConflict
	InternalServerError ErrorCode = http.StatusInternalServerError
)

// Code ...
func (e ErrorCode) Code() int32 {
	return int32(e)
}
