package httperror

import (
	"fmt"
	"net/http"
	"strings"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/multierr"
)

type ErrorMap struct {
	Errors []interface{} `json:"errors"`
}

type HttpErrorHandler struct {
	StatusCode int
	Content    *ErrorMap
}

func (e HttpErrorHandler) Error() string {
	return ""
}

func NewNotFoundError(entity string) *HttpErrorHandler {
	return &HttpErrorHandler{
		StatusCode: http.StatusNotFound,
		Content:    MapErrors(fmt.Errorf("%v not found", entity)),
	}
}

func NewBadRequestError(err error) *HttpErrorHandler {
	return &HttpErrorHandler{
		StatusCode: http.StatusBadRequest,
		Content:    MapErrors(err),
	}
}

func NewResourceAlreadyExistsError(entity string) *HttpErrorHandler {
	return &HttpErrorHandler{
		StatusCode: http.StatusConflict,
		Content:    MapErrors(fmt.Errorf("%v is already registered", entity)),
	}
}

func NewConflictError(msg string) *HttpErrorHandler {
	return &HttpErrorHandler{
		StatusCode: http.StatusConflict,
		Content:    MapErrors(fmt.Errorf("%s", msg)),
	}
}

func NewInternalServerError(err error) *HttpErrorHandler {
	return &HttpErrorHandler{
		StatusCode: http.StatusInternalServerError,
		Content:    MapErrors(err),
	}
}

func NewError(err error) *HttpErrorHandler {
	if err.Error() == "multiple processes tried to update a resource at the same time" {
		return NewConflictError(err.Error())
	}

	return NewInternalServerError(err)
}

func MapErrors(err error) *ErrorMap {
	m := &ErrorMap{}

	if errs, ok := err.(ozzo.Errors); ok {
		for _, em := range strings.Split(fmt.Sprintf("%s", errs), ";") {
			m.Errors = append(m.Errors, em)
		}
	} else {
		for _, e := range multierr.Errors(err) {
			m.Errors = append(m.Errors, e.Error())
		}
	}

	return m
}
