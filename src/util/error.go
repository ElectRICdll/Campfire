package util

import (
	"fmt"
)

func NewExternalError(msg string) error {
	return ExternalError{
		Message: msg,
	}
}

type ExternalError struct {
	Message string
}

func (err ExternalError) Error() string {
	return fmt.Sprintf("External Error: %s", err.Message)
}
