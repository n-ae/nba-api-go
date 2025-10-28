package models

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidResponse = errors.New("invalid response format")
	ErrRateLimited     = errors.New("rate limited")
	ErrNotFound        = errors.New("resource not found")
	ErrInvalidRequest  = errors.New("invalid request parameters")
	ErrTimeout         = errors.New("request timeout")
	ErrUnauthorized    = errors.New("unauthorized")
)

type APIError struct {
	StatusCode int
	Message    string
	URL        string
	Err        error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("API error (status %d, url %s): %s: %v", e.StatusCode, e.URL, e.Message, e.Err)
	}
	return fmt.Sprintf("API error (status %d, url %s): %s", e.StatusCode, e.URL, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

func NewAPIError(statusCode int, url, message string, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		URL:        url,
		Message:    message,
		Err:        err,
	}
}

func HTTPStatusToError(statusCode int, url string) error {
	switch statusCode {
	case http.StatusNotFound:
		return NewAPIError(statusCode, url, "resource not found", ErrNotFound)
	case http.StatusUnauthorized, http.StatusForbidden:
		return NewAPIError(statusCode, url, "unauthorized", ErrUnauthorized)
	case http.StatusTooManyRequests:
		return NewAPIError(statusCode, url, "rate limited", ErrRateLimited)
	case http.StatusBadRequest:
		return NewAPIError(statusCode, url, "invalid request", ErrInvalidRequest)
	case http.StatusGatewayTimeout, http.StatusRequestTimeout:
		return NewAPIError(statusCode, url, "request timeout", ErrTimeout)
	default:
		if statusCode >= 500 {
			return NewAPIError(statusCode, url, "server error", fmt.Errorf("HTTP %d", statusCode))
		}
		if statusCode >= 400 {
			return NewAPIError(statusCode, url, "client error", fmt.Errorf("HTTP %d", statusCode))
		}
		return nil
	}
}
