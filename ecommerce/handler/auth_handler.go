package handler

import (
	"ecommerce/helpers"
	"ecommerce/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) AuthHandler {
	return AuthHandler{
		db: db,
	}
}

// Login godoc
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	var user models.User

	err = h.db.Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		c.JSON(401, helpers.NewErrorResponse[any](401, "Email not found"))
		return
	}

	// if req.Password != user.Password {
	// 	c.JSON(401, helpers.NewErrorResponse[any](401, "Invalid password"))
	// 	return
	// }

	// check password
	_, err = helpers.VerifyPassword(req.Password, user.Password)

	if err != nil {
		c.JSON(401, helpers.NewErrorResponse[any](401, "Invalid password"))
		return
	}

	// generate token

	exp := time.Now().Add(time.Hour * 24)

	token, err := helpers.GenerateToken(user, &exp)

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to generate token"))
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

// Register godoc
func (h *AuthHandler) Register(c *gin.Context) {

	var req models.RegisterUserRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	hashPassword, err := helpers.HashPassword(req.Password)

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to create user"))
		return
	}

	otpCode := helpers.GenerateOTP(6)

	// create user
	user := models.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashPassword,
		OTPCode:  &otpCode,
		RoleID:   2,
	}

	// save user to database
	err = h.db.Create(&user).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to create user"))
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req models.VerifyOTPRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, helpers.NewValidationResponse[any](400, "Bad request", err))
		return
	}

	var user *models.User

	err = h.db.Where("email = ?", req.Email).First(&user).Error

	if err != nil {
		c.JSON(401, helpers.NewErrorResponse[any](401, "Email not found"))
		return
	}

	// check OTP
	if req.OTP != *user.OTPCode {
		c.JSON(401, helpers.NewErrorResponse[any](401, "Invalid OTP"))
		return
	}

	// update user
	err = h.db.Model(&user).Update("is_verified", true).Error

	if err != nil {
		c.JSON(500, helpers.NewErrorResponse[any](500, "Failed to verify OTP"))
		return
	}

	c.JSON(200, gin.H{
		"message": "OTP verified",
	})
}
