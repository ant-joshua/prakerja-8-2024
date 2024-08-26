package models

type User struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	IsVerified bool    `json:"is_verified" gorm:"default:false"`
	OTPCode    *string `json:"otp_code" gorm:"column:otp_code"` // 6 digit code
}

type RegisterUserRequest struct {
	Name     string `json:"name" name:"name" binding:"required"`
	Email    string `json:"email" name:"email" binding:"required"`
	Password string `json:"password" name:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" name:"email" binding:"required"`
	Password string `json:"password" name:"password" binding:"required"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" name:"email" binding:"required"`
	OTP   string `json:"otp" name:"otp" binding:"required"`
}
