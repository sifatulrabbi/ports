package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/models"
	"github.com/sifatulrabbi/ports/pkg/services"
)

// Auth decorator
// This will first validate the request and will forward it to the next function
func AuthGuard(next func(w http.ResponseWriter, r *http.Request, session *models.Session)) func(w http.ResponseWriter, r *http.Request) {
	res := Response{}

	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		// Confirm the access token
		if bearerToken == "" {
			res.Message = "no access token found"
			res.Unauthorized(w)
			return
		}
		splitToken := strings.Split(bearerToken, " ")
		if splitToken[0] != "Bearer" || splitToken[1] == "" {
			res.Message = "invalid access token"
			res.Unauthorized(w)
			return
		}
		// Verify the access token
		accessToken := splitToken[1]

		claims := models.AccessTokenClaims{}
		token, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configs.Globals.JWT_SECRET), nil
		})

		if err != nil {
			res.Message = err.Error()
			res.Forbidden(w)
			return
		}

		session := models.Session{}
		if err = token.Claims.Valid(); err == nil && token.Valid {
			// Validate the auth token.
			authClaims := models.AuthTokenClaims{}
			authToken, err := jwt.ParseWithClaims(claims.AuthToken, &authClaims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected singing method: %v", token.Header["alg"])
				}
				session, err = services.FindSessionById(claims.ID)
				if err != nil {
					return []byte(""), err
				}
				return []byte(session.RefreshToken), nil
			})
			if err != nil {
				res.Message = err.Error()
				res.BadRequest(w)
				return
			}
			if err = authToken.Claims.Valid(); err != nil {
				res.Message = err.Error()
				res.BadRequest(w)
				return
			}
			next(w, r, &session)
		} else {
			res.Message = err.Error()
			res.Forbidden(w)
		}
	}
}
