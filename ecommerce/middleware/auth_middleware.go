package middleware

import (
	"ecommerce/helpers"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		getHeader := c.GetHeader("Authorization")

		// check if authorization header is empty
		if getHeader == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		// get token from authorization header
		authHeader := strings.Split(getHeader, " ")

		// Authorization: Bearer token

		// ["Bearer", "123123123123123"]

		if len(authHeader) != 2 {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		bearerToken := authHeader[1]
		// check if token is valid

		token, err := helpers.ValidateToken(bearerToken)

		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized", // token is invalid
			})

			c.Abort()
			return
		}

		tokenClaims := token.Claims.(jwt.MapClaims)

		fmt.Printf("Token claims: %+v\n", tokenClaims)

		c.Set("user", tokenClaims)

		// if token is valid, continue
		c.Next()
	}
}
