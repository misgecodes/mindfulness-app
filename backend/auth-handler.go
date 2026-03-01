package main

import (
	// "encoding/json"
	"fmt"
	"mindfulness-app/database"
	"net/http"

	"time"

	// "net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// for refresh tokens

func Register(c *gin.Context) {
	user := UserRegister{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(400, "Invalid input format")
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT 1 FROM users WHERE email =$1",
		user.Email).Scan(&exists)
	if err != nil {
		fmt.Println("Database error:", err)
	}

	if exists {
		c.AbortWithStatusJSON(400, gin.H{"error": "Email already registered"})
		return
	}
	var user_id int

	err = database.DB.QueryRow("insert into users(email, password) values ($1, $2) RETURNING id",
		user.Email,
		user.Password).Scan(&user_id)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create user", "details": err.Error()})
	}

	accessToken, err := GenerateAccessToken(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Access Token", "details": err.Error()})
		return
	}

	refreshToken, err := GenerateRefreshToken(user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Refresh Token", "details": err.Error()})
		return
	}

	expiresAt := time.Now().Add(RefreshTokenTTL)

	_, err = database.DB.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)",
		user_id,
		refreshToken,
		expiresAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Save Refresh Token"})
		return
	}

	c.IndentedJSON(200, LoginResponse{
		Message:      "Registration successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken})

}

func Login(c *gin.Context) {

	login_request := UserLogin{}
	if err := c.ShouldBindJSON(&login_request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Login  data"})
		return
	}
	user := User{}

	err := database.DB.QueryRow("SELECT id, email, password FROM users WHERE email =$1 AND password = $2",
		login_request.Email, login_request.Password).Scan(&user.ID, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if err != nil {
		fmt.Println("Database error:", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Database error"})
		return
	}

	// generate acess token and refresh token here and send to client
	accessToken, err := GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Access Token", "details": err.Error()})
		return
	}

	refreshToken, err := GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Generate Refresh Token", "details": err.Error()})
		return
	}

	expiresAt := time.Now().Add(RefreshTokenTTL)

	_, err = database.DB.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)",
		user.ID,
		refreshToken,
		expiresAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Save Refresh Token"})
		return
	}

	c.IndentedJSON(200, LoginResponse{
		Message:      "Login successful",
		AccessToken:  accessToken,
		RefreshToken: refreshToken})
}
