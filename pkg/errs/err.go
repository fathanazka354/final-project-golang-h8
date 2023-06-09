package errs

import "net/http"

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type MessageErrData struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *MessageErrData) Message() string {
	return e.ErrMessage
}
func (e *MessageErrData) Status() int {
	return e.ErrStatus
}
func (e *MessageErrData) Error() string {
	return e.ErrError
}

func NewUnauthorizedError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "NOT_AUTHORIZED",
	}
}
func NewUnauthenticatedError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
	}
}
func NewNotFoundError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusNoContent,
		ErrError:   "NOT_CONTENT",
	}
}
func NewNotContent(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusNoContent,
		ErrError:   "NOT_CONTENT",
	}
}
func NewBadRequest(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}
func NewInternalServerError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}
func NewUnprocessibleEntityError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}
