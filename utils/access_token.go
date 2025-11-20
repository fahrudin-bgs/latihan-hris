package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(userID uint) string {
	// buat access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 1 hari
	})

	// token
	accessTokenString, _ := accessToken.SignedString([]byte(os.Getenv("SECRET")))

	return accessTokenString
}
