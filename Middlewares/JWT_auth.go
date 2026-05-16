package Middlewares

import (
	"net/http"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("hellogolang")

func AuthMiddleware(requiredRole string) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Headers")
		fmt.Println("auth header", authHeader)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No token provided",
			})
			c.Abort()
			return
		}

		// Remove "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		fmt.Println("the token is ", token)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid claims",
			})
			c.Abort()
			return
		}

		role := claims["role"].(string)
		fmt.Println("the role is ", role)

		if role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied",
			})
			c.Abort()
			return
		}

		// Store values in context
		c.Set("username", claims["username"])
		c.Set("role", role)

		c.Next()
	}
}
