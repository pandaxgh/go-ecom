package errs

import "net/http"

func NewUnauthorizedError(message string) *HttpError {
	return &HttpError{
		Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusUnauthorized)),
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *HttpError {
	return &HttpError{
		Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusForbidden)),
		Message: message,
		Status:  http.StatusForbidden,
	}
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusBadRequest)),
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusNotFound)),
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func NewInternalServerError() *HttpError {
	return &HttpError{
		Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusInternalServerError)),
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
	}
}
