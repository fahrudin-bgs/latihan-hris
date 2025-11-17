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
		"exp": time.Now().Add(time.Minute * 30).Unix(), // 30 menit
	})

	// token
	accessTokenString, _ := accessToken.SignedString([]byte(os.Getenv("SECRET")))

	return accessTokenString
}
