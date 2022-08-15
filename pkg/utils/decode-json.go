package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ErrResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ErrResp) Error() string {
	return e.Message
}

func DecodeJson(w http.ResponseWriter, r *http.Request, v interface{}) (errResp *ErrResp) {
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		errResp.Message = "Cannot accept request body"
		errResp.Status = 400
		return
	}
	w.Header().Add("Content-Type", "application/json")
	// Reading body with a limited bytes size.
	// This will throw an error if the size of the body bigger than what is specified.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// JSON decoder to decode the body.
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // Disallowing *unknown* unspecified fields.
	err := dec.Decode(&v)
	if err != nil {
		var (
			syntaxErr        *json.SyntaxError
			unmarshalTypeErr *json.UnmarshalTypeError
		)
		errResp.Status = 400
		switch {
		case errors.As(err, &syntaxErr):
			errResp.Message = "Request body contains malformed data"

		case errors.As(err, &unmarshalTypeErr):
			errResp.Message = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeErr.Field, unmarshalTypeErr.Offset)

		case errors.Is(err, io.EOF):
			errResp.Message = "Request body must not be empty"

		case err.Error() == "http: request body too large":
			errResp.Message = "HTTP body must not be larger than 1MB"

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errResp.Message = fmt.Sprintf("Request body contains unknown field %s", fieldName)

		default:
			errResp.Message = "Unable to read request body"
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		errResp.Message = "Request body must only contain a single JSON object"
		errResp.Status = 400
	}

	return
}
