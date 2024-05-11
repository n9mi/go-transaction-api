package exception

import (
	"strings"
)

type HttpError struct {
	Code     int
	Messages []string
}

func (e *HttpError) Error() string {
	return strings.Join(e.Messages, ",")
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{Code: code, Messages: []string{message}}
}
