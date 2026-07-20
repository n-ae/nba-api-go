package models

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidResponse  = errors.New("invalid response format")
	ErrRateLimited      = errors.New("rate limited")
	ErrNotFound         = errors.New("resource not found")
	ErrInvalidRequest   = errors.New("invalid request parameters")
	ErrTimeout          = errors.New("request timeout")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrResponseTooLarge = errors.New("response body exceeds configured maximum size")
)

// maxErrorBodyBytes bounds how much of a failing response body APIError
// keeps for diagnostics - enough to see the shape of an NBA.com or proxy
// error page without risking an oversized error message.
const maxErrorBodyBytes = 2048

type APIError struct {
	StatusCode int
	Message    string
	URL        string
	// Body is a truncated (at most maxErrorBodyBytes) copy of the
	// response body, kept for diagnostics; the full body is discarded.
	Body string
	Err  error
}

func (e *APIError) Error() string {
	if e.Err != nil {
		if e.Body != "" {
			return fmt.Sprintf("API error (status %d, url %s): %s: %v (body: %q)", e.StatusCode, e.URL, e.Message, e.Err, e.Body)
		}
		return fmt.Sprintf("API error (status %d, url %s): %s: %v", e.StatusCode, e.URL, e.Message, e.Err)
	}
	if e.Body != "" {
		return fmt.Sprintf("API error (status %d, url %s): %s (body: %q)", e.StatusCode, e.URL, e.Message, e.Body)
	}
	return fmt.Sprintf("API error (status %d, url %s): %s", e.StatusCode, e.URL, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// NewAPIError builds an APIError, truncating body to at most
// maxErrorBodyBytes before storing it.
func NewAPIError(statusCode int, url, message string, body []byte, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		URL:        url,
		Message:    message,
		Body:       truncateBody(body),
		Err:        err,
	}
}

func truncateBody(body []byte) string {
	if len(body) > maxErrorBodyBytes {
		return string(body[:maxErrorBodyBytes])
	}
	return string(body)
}

func HTTPStatusToError(statusCode int, url string, body []byte) error {
	switch statusCode {
	case http.StatusNotFound:
		return NewAPIError(statusCode, url, "resource not found", body, ErrNotFound)
	case http.StatusUnauthorized, http.StatusForbidden:
		return NewAPIError(statusCode, url, "unauthorized", body, ErrUnauthorized)
	case http.StatusTooManyRequests:
		return NewAPIError(statusCode, url, "rate limited", body, ErrRateLimited)
	case http.StatusBadRequest:
		return NewAPIError(statusCode, url, "invalid request", body, ErrInvalidRequest)
	case http.StatusGatewayTimeout, http.StatusRequestTimeout:
		return NewAPIError(statusCode, url, "request timeout", body, ErrTimeout)
	default:
		if statusCode >= 500 {
			return NewAPIError(statusCode, url, "server error", body, fmt.Errorf("HTTP %d", statusCode))
		}
		if statusCode >= 400 {
			return NewAPIError(statusCode, url, "client error", body, fmt.Errorf("HTTP %d", statusCode))
		}
		return nil
	}
}
