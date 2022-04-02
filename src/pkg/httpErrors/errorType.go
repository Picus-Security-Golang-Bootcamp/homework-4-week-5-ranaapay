package httpErrors

import (
	"errors"
	"net/http"
	"strings"
)

type ErrorType struct {
	ErrCode    int    `json:"errCode"`
	ErrMessage string `json:"errMessage"`
}

func NewErrorType(code int, message string) ErrorType {
	return ErrorType{
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NewInternalServerError(err error) ErrorType {
	return ErrorType{
		ErrCode:    http.StatusInternalServerError,
		ErrMessage: err.Error(),
	}
}

func ParseError(err error) ErrorType {
	switch {
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewErrorType(http.StatusBadRequest, err.Error())
	case strings.Contains(err.Error(), "is required"):
		return NewErrorType(http.StatusBadRequest, err.Error())
	case strings.Contains(err.Error(), "Not Found"):
		return NewErrorType(http.StatusNotFound, err.Error())
	case errors.Is(err, ContentTypeError):
		return NewErrorType(http.StatusBadRequest, err.Error())
	case errors.Is(err, ConvertIdError):
		return NewErrorType(http.StatusBadRequest, err.Error())
	case errors.Is(err, CannotDecodeError):
		return NewErrorType(http.StatusBadRequest, err.Error())
	default:
		return NewInternalServerError(err)
	}
}
