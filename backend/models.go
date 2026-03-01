package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Topic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type Content struct {
	ID      int    `json:"id"`
	TopicID int    `json:"topic_id"`
	Title   string `json:"title"`
	URI     string `json:"uri"`
	// Duration int    `json:"duration"`
}

type claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserTopic struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	TopicID   string `json:"topic_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsActive  bool   `json:"is_active"`
}

type User struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
}

type RefreshToken struct {
	ID        string    `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"refresh_token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	Revoked   bool      `json:"revoked"`
}
