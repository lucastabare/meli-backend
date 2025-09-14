package util

import "net/http"

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e AppError) Error() string { return e.Message }

var (
	ErrNotFound   = AppError{Code: http.StatusNotFound, Message: "resource not found"}
	ErrBadRequest = AppError{Code: http.StatusBadRequest, Message: "bad request"}
)
