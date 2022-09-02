package utils

import (
	"encoding/json"
	"net/http"
)

// Parse the body into provided format.
func BodyParser(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	return err
}
