package utils

import (
	"net/http"
	"strings"
)

func AuthGuard(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	res := CustomResponse{}

	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		// Confirm the access token
		if accessToken == "" {
			res.Message = "no access token found"
			res.Unauthorized(w)
			return
		}
		splitToken := strings.Split(accessToken, " ")
		if splitToken[0] != "Bearer" || splitToken[1] == "" {
			res.Message = "invalid access token"
			res.Unauthorized(w)
			return
		}
		// Verify the access token
	}
}
