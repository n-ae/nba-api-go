package models

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Data       T
	StatusCode int
	URL        string
	Headers    http.Header
}

func NewResponse[T any](data T, statusCode int, url string, headers http.Header) *Response[T] {
	return &Response[T]{
		Data:       data,
		StatusCode: statusCode,
		URL:        url,
		Headers:    headers,
	}
}

func (r *Response[T]) JSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

type RawResponse struct {
	Body       []byte
	StatusCode int
	URL        string
	Headers    http.Header
}

func NewRawResponse(body []byte, statusCode int, url string, headers http.Header) *RawResponse {
	return &RawResponse{
		Body:       body,
		StatusCode: statusCode,
		URL:        url,
		Headers:    headers,
	}
}

func (r *RawResponse) Unmarshal(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}
