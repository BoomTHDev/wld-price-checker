package custom

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("AppError: %s - %s", e.Message, e.Err)
	}
	return fmt.Sprintf("AppError: %s", e.Message)
}

func NewError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func ErrNotFound(message string, err error) *AppError {
	if message == "" {
		message = "The request resource was not found."
	}
	return NewError(fiber.StatusNotFound, message, err)
}

func ErrInvalidInput(message string, err error) *AppError {
	if message == "" {
		message = "Invalid input provided."
	}
	return NewError(fiber.StatusBadRequest, message, err)
}

func ErrUnauthorized(message string, err error) *AppError {
	if message == "" {
		message = "Unauthorized access."
	}
	return NewError(fiber.StatusUnauthorized, message, err)
}

func ErrForbidden(message string, err error) *AppError {
	if message == "" {
		message = "Access to this resource is forbidden."
	}
	return NewError(fiber.StatusForbidden, message, err)
}

func ErrIntervalServer(message string, err error) *AppError {
	if message == "" {
		message = "An unexcepted internal server error occurred."
	}
	return NewError(fiber.StatusInternalServerError, message, err)
}

func ErrConflict(message string, err error) *AppError {
	if message == "" {
		message = "A conflict occurred with the current state of the resource."
	}
	return NewError(fiber.StatusConflict, message, err)
}

func ErrTooManyRequests(message string, err error) *AppError {
	if message == "" {
		message = "Too many requests."
	}
	return NewError(fiber.StatusTooManyRequests, message, err)
}
