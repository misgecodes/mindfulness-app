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
	// database.MigrateUsersTable(database.DB)
	// database.MigrateTopicsTable(database.DB)
	// database.MigrateContentsTable(database.DB)
	database.MigrateRefreshTokensTable(database.DB)
	router := gin.Default()
	router.GET("/topics", getTopics)
	router.GET("/contents/:topic-id", getContents)
	router.POST("/add-topic", addTopic)
	router.POST("/add-content", addContent)
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"status": "OK"})
	})
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/refresh-token", RefereshAccessToken)
	router.GET("/users", getUsers)
	router.Run(":8080")
}
