package errs

import (
	"encoding/json"
	"net/http"
	"strings"
)

type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) WithMessage(message string) *HttpError {
	return &HttpError{
		Code:    e.Code,
		Message: message,
		Status:  e.Status,
	}
}

func ErrorHandler(next AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if e, ok := err.(*HttpError); ok {
				w.WriteHeader(e.Status)
				json.NewEncoder(w).Encode(e)
				return
			}
			internalError := &HttpError{
				Code:    MakeAllUppercaseWithUnderscores(http.StatusText(http.StatusInternalServerError)),
				Message: http.StatusText(http.StatusInternalServerError),
				Status:  http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(internalError)
		}
	}
}

func MakeAllUppercaseWithUnderscores(str string) string {
	return strings.ToUpper(strings.ReplaceAll(str, " ", "_"))
}
