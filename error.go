package entity

import (
	"fmt"
)

type ExternalError struct {
	Message string
}

func (err ExternalError) Error() string {
	return fmt.Sprintf("External Error: %s", err.Message)
}
