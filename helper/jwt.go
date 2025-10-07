package helper

import (
	"first-rest-api-go/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.GetEnv("JWT_SECRET_KEY", "JWT_SECRET_KEY"))

type UserClaims struct {
	UserID uint    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId uint, email string) (string, error) {
	// Mengatur waktu kedaluwarsa token, di sini kita set 60 menit dari waktu sekarang
	expirationTime := time.Now().Add(60 * time.Minute)

	// Create the JWT claims, which includes the user ID and expiry time
	// Subject = userId with email
	claims := &UserClaims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}