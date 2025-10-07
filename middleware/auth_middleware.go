package middleware

import (
	"net/http"
	"first-rest-api-go/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET_KEY", "JWT_SECRET_KEY"))

func AuthMiddleware() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		authTokenHeader := c.GetHeader("Authorization")

		if authTokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		authTokenHeader = strings.TrimPrefix(authTokenHeader, "Bearer ")

		claims := jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(authTokenHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}