package errors

import (
    "errors"
    "fmt"
    "net/http"
)

type AppError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Err     error  `json:"-"`
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Err
}

func NewBadRequestError(message string, err error) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: message,
        Err:     err,
    }
}

func NewNotFoundError(message string, err error) *AppError {
    return &AppError{
        Code:    http.StatusNotFound,
        Message: message,
        Err:     err,
    }
}

func NewInternalServerError(message string, err error) *AppError {
    return &AppError{
        Code:    http.StatusInternalServerError,
        Message: message,
        Err:     err,
    }
}

func NewMethodNotAllowedError() *AppError {
    return &AppError{
        Code:    http.StatusMethodNotAllowed,
        Message: "Method not allowed",
    }
}

var (
    ErrInvalidLatitude     = errors.New("latitude must be between -90 and 90")
    ErrInvalidLongitude    = errors.New("longitude must be between -180 and 180")
    ErrMissingAffectedArea = errors.New("affected area is required")
    ErrMissingActions      = errors.New("actions are required")
    ErrNoDischargeData     = errors.New("no discharge data available")
)
