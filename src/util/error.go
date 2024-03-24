package util

import (
	"campfire/log"
	"fmt"
	"runtime"
	"strings"
)

func NewExternalError(msg string) error {
	n, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(n)
	log.Debug(f.Name())
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

func NewErrorGroup(errs ...error) *ExternalErrorGroup {
	return &ExternalErrorGroup{Errors: errs}
}

type ExternalErrorGroup struct {
	Errors []error
}

func (e *ExternalErrorGroup) AddError(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *ExternalErrorGroup) Error() string {
	var b strings.Builder
	for i, err := range e.Errors {
		if i > 0 {
			b.WriteString("; ")
		}
		b.WriteString(err.Error())
	}
	return b.String()
}
