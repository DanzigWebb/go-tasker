package models

type ApiError struct {
	Error   bool   `json:"error" default:"true"`
	Message string `json:"message"`
}

func DefaultError(msg string) ApiError {
	return ApiError{
		Error:   true,
		Message: msg,
	}
}
