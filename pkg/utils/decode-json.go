package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type ErrResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func DecodeJson(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	ct := r.Header.Get("Content-Type")
	log.Println(ct)

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
			msg              string
		)

		switch {
		case errors.As(err, &syntaxErr):
			msg = "Request body contains malformed data"

		case errors.As(err, &unmarshalTypeErr):
			msg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeErr.Field, unmarshalTypeErr.Offset)

		case errors.Is(err, io.EOF):
			msg = "Request body must not be empty"

		case err.Error() == "http: request body too large":
			msg = "HTTP body must not be larger than 1MB"

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg = fmt.Sprintf("Request body contains unknown field %s", fieldName)

		default:
			msg = "Unable to read request body"
		}

		rb, _ := json.Marshal(&ErrResp{Status: 400, Message: msg})
		if _, err := w.Write(rb); err != nil {
			log.Fatal(err)
		}
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		rb, _ := json.Marshal(&ErrResp{Status: 400, Message: msg})
		if _, err := w.Write(rb); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
