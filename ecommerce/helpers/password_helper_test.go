package helpers

import "testing"

func TestGenerateOTP(t *testing.T) {
	otp := GenerateOTP(6)

	if len(otp) != 6 {
		t.Errorf("OTP length is not 6")
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "password"
	hashedPassword, _ := HashPassword(password)

	result, _ := VerifyPassword(password, hashedPassword)

	if result != true {
		t.Errorf("Password verification failed")
	}
}
