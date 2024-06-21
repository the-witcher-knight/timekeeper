package httpio

import (
	"encoding/json"
	"io"
	"net/http"
)

// BindJSON decodes JSON data from the given io.Reader into a provided type
func BindJSON[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		if err != io.EOF {
			return *new(T), err
		}
	}

	return body, nil
}

func DecodeValidate[T Validator](w http.ResponseWriter, r *http.Request) (T, error) {
	req, err := BindJSON[T](w, r)
	if err != nil {
		return *new(T), err
	}

	if err := req.Valid(r.Context()); err != nil {
		return *new(T), err
	}

	return req, nil
}
