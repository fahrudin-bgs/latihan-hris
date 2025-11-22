package middleware

import (
	"fmt"
	"latihan-hris/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Missing Authorization header")
			ctx.Abort()
			return
		}

		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	ctx.JSON(401, gin.H{"error": "Invalid Authorization format"})
		// 	ctx.Abort()
		// 	return
		// }

		tokenString := authHeader

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid claims")
			ctx.Abort()
			return
		}

		userID := fmt.Sprintf("%v", claims["sub"])

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
