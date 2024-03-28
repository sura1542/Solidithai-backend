package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				return nil, fmt.Errorf("invalid toke: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])

		} else {
			c.JSON(http.StatusOK, gin.H{"status": "forbidden", "message": err.Error()})
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "forbidden", "message": err.Error()})
		}
		c.Next()
	}
}
