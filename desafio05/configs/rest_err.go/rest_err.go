package resterr

import (
	"net/http"

	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewBadRequestError(message string, causes ...Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  nil,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
		Causes:  nil,
	}
}

func ConvertError(internalerrors *internalerrors.InternalError) *RestErr {
	switch internalerrors.Err {
	case "bad_request":
		return NewBadRequestError(internalerrors.Error())
	case "not_found":
		return NewNotFoundError(internalerrors.Error())
	default:
		return NewInternalServerError(internalerrors.Error())
	}
}
