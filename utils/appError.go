package main

import (
	"fmt"
)

// AppError defines a custom error type
type AppError struct {
	Message       string `json:"message"`
	StatusCode    int    `json:"statusCode"`
	Status        string `json:"status"`
	IsOperational bool   `json:"isOperational"`
}

func NewAppError(message string, statusCode int) *AppError {
	status := "Error"
	if statusCode >= 400 && statusCode < 500 {
		status = "Fail"
	}
	return &AppError{
		Message:       message,
		StatusCode:    statusCode,
		Status:        status,
		IsOperational: true,
	}
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s (status %d)", e.Status, e.Message, e.StatusCode)
}
