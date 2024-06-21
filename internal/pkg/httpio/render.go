package httpio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/the-witcher-knight/timekeeper/internal/pkg/tracing"
)

// WriteJSON writes a JSON response to the provided http.ResponseWriter
func WriteJSON[T any](w http.ResponseWriter, r *http.Request, data Response[T]) {
	// Prepare HTTP response headers
	w.Header().Set("Content-Type", "application/json")
	for key, val := range data.Headers {
		w.Header().Set(key, val)
	}

	// Update the HTTP status code
	w.WriteHeader(data.Status)

	// Encode the response body to JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(data.Body); err != nil {
		tracing.FromCtx(r.Context()).Error(context.Background(),
			err, "error when encode JSON response", true)
	}
}

// Response represents an HTTP response structure with a generic body 'T'.
// It includes the HTTP status code, headers, and the response body.
// The type 'T' can be any type, making it flexible for different response structures.
type Response[T any] struct {
	Status  int
	Headers map[string]string
	Body    T
}

// Error represents an application-specific error structure.
// It implements the error interface, allowing it to be used as an error type.
type Error struct {
	Status int
	Code   string
	Desc   string
}

func (e Error) Error() string {
	return fmt.Sprintf(`{"code":"%s","desc":"%s"}`, e.Code, e.Desc)
}

// Message represents a generic message structure used for communication.
// Code is message code, Desc is message description
type Message struct {
	Code string `json:"code"`
	Desc string `json:"desc,omitempty"`
}
