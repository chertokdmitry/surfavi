package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json: "message"`
	Status  int    `json: "status"`
	Error   string `json:"error"`
}

// return new error
func NewError(msg string) error {
	return errors.New(msg)
}

// bad request error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

}

// not found error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}

}

// internal server error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}

}
