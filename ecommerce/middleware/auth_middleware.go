package middleware

import (
	"ecommerce/helpers"
	"ecommerce/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
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

		userID := tokenClaims["id"].(float64)

		// check from database
		var user models.User

		err = db.Where("id = ?", userID).First(&user).Error

		if err != nil {
			c.JSON(401, gin.H{
				"message": "Unauthorized", // user not found
			})

			c.Abort()
			return
		}

		if user.Email != tokenClaims["email"] {

			c.JSON(401, gin.H{
				"message": "Unauthorized", // email not match
			})

			c.Abort()
			return
		}

		// check if verify email

		if !user.IsVerified {
			c.JSON(401, gin.H{
				"message": "Unauthorized", // email not verified
			})

			c.Abort()
			return
		}

		// check if the user is already verify their email
		// check from database

		c.Set("user", tokenClaims)

		// if token is valid, continue
		c.Next()
	}
}
