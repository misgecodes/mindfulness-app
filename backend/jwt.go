package main

import (
	"crypto/rand"
	"encoding/base64"
	"mindfulness-app/database"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	// "math/rand"
)

var (
	JWTSecret       = []byte(os.Getenv("JWTSecret"))
	AccessTokenTTL  = 15 * time.Minute
	RefreshTokenTTL = 7 * 24 * time.Hour
)

func GenerateAccessToken(userID int) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)

}

func GenerateRefreshToken(userID int) (string, error) {

	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func RefereshAccessToken(c *gin.Context) {
	// get refresh token from request body
	var token struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refersh Token Required "})
		return
	}
	var storedToken RefreshToken
	err := database.DB.QueryRow("SELECT id, user_id, token, expires_at, is_revoked FROM refresh_tokens WHERE token = $1",
		token.RefreshToken).Scan(&storedToken.ID,
		&storedToken.UserID,
		&storedToken.Token,
		&storedToken.ExpiresAt,
		&storedToken.Revoked)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Refresh Token", "details": err.Error()})
		return
	}

	if storedToken.Revoked || time.Now().After(storedToken.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh Token Expired or Revoked"})
		return
	}

	accessToken, err := GenerateAccessToken(storedToken.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Access Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})

}

func verifyOTP(otp string) bool {
	return otp == "123456" // Placeholder for demonstration
}

// func generateOTP(length int) string {
//     digits := "0123456789"
//     otp := make([]byte, length)
//     for i := range otp {
//         otp[i] = digits[rand.Intn(len(digits))]
//     }
//     return string(otp)
// }
