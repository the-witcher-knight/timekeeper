package respond

import (
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/httpio"
)

// Created for write (201) response
func Created[T any](val T) Response[T] {
	return Response[T]{
		Response: httpio.Response[T]{
			Status: http.StatusCreated,
			Body:   val,
		},
	}
}

// OK for write (200) response
func OK[T any](val T) Response[T] {
	return Response[T]{
		Response: httpio.Response[T]{
			Status: http.StatusOK,
			Body:   val,
		},
	}
}

// NoContent for write (204) response
func NoContent[T any](val T) Response[T] {
	return Response[T]{
		Response: httpio.Response[T]{
			Status: http.StatusNoContent,
			Body:   val,
		},
	}
}

type Response[T any] struct {
	httpio.Response[T]
}

func (resp Response[T]) WithHeaders(headers map[string]string) {
	for key, val := range headers {
		resp.Headers[key] = val
	}
}

func (resp Response[T]) Write(w http.ResponseWriter, r *http.Request) {
	httpio.WriteJSON(w, r, httpio.Response[T]{
		Status: resp.Status,
		Body:   resp.Body,
	})
}
