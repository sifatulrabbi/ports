package services

import (
	"crypto/sha256"
	"net"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/sifatulrabbi/ports/pkg/configs"
	"github.com/sifatulrabbi/ports/pkg/models"
)

func CreateSession(r *http.Request, u *models.User) (refreshToken string, accessToken string, err error) {
	// get the ip address of the client
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	// create a refresh token
	hasher := sha256.New()
	hasher.Write([]byte(uuid.NewString()))
	token := string(hasher.Sum(nil))
	rClaims := models.RefreshTokenClaims{
		UserID:       u.ID,
		Username:     u.Username,
		Iat:          time.Now().UnixMilli(),
		RefreshToken: token,
		IP:           ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 365 * time.Hour)),
		},
	}
	rJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	refreshToken, err = rJwtToken.SignedString([]byte(configs.Globals.JWT_SECRET))
	if err != nil {
		return "", "", err
	}

	// create jwt token for the client.
	aClaims := models.AuthTokenClaims{
		UserID:   u.ID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			Issuer:    "ports-app",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	accessToken, err = jwtToken.SignedString([]byte(configs.Globals.JWT_SECRET))
	if err != nil {
		return "", "", err
	}
	return
}
