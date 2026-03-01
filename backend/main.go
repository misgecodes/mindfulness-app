package main

import (
	"mindfulness-app/database"
	// "log"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10/translations/id"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	database.ConnectDatabase()
	database.MigrateUsersTable(database.DB)
	// database.MigrateTopicsTable(database.DB)
	// database.MigrateContentsTable(database.DB)
	// database.MigrateOTPTable(database.DB)
	// database.MigrateRefreshTokensTable(database.DB)
	router := gin.Default()
	router.GET("/topics", getTopics)
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/refresh-token", RefereshAccessToken)

	// Protected routes
	auth := router.Group("/api")
	auth.Use(JWTAuth())
	{
		auth.GET("/contents/:topic-id", getContents)
		auth.POST("/add-topic", addTopic)
		auth.POST("/add-content", addContent)
		auth.GET("/users", getUsers)
	}
	// router.GET("/contents/:topic-id", getContents)
	// router.POST("/add-topic", addTopic)
	// router.POST("/add-content", addContent)
	// router.GET("/users", getUsers)
	router.Run(":8080")

}
