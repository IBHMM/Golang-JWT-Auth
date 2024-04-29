package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.JSON(401, gin.H{"error": err, "message": "Unauthorized"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			var secret string = os.Getenv("SECRET")
			return []byte(secret), nil
		})
		if err != nil {
			ctx.JSON(401, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := claims["exp"].(float64)
			if int64(exp) < time.Now().Unix() {
				ctx.JSON(401, gin.H{"message": "Token has expired"})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(401, gin.H{"message": "Invalid token claims"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
