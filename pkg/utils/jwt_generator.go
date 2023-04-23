package utils

import (
	models_authentication "app/app/models/service_authentication"

	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// สร้าง Token และเก็บไว้ใน
func GenerateAccessTokenUser(user models_authentication.Req_login) (string, error) {
	// Set secret key from .env file.
	secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["id"] = user.Username

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
