package main

import (
	"mindfulness-app/database"
	// "log"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10/translations/id"
)

// func getTopics(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, topics)
// }

func main() {
	gin.SetMode(gin.ReleaseMode)
	database.ConnectDatabase()
	database.MigrateUsersTable(database.DB)
	router := gin.Default()
	router.GET("/topics", getTopics)
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"status": "OK"})
	})
	router.POST("/add-user", addUser)
	router.GET("/users", getUsers)
	// router.POST("/user-topics", createUserTopic)
	// router.PUT("/user-topics/:id", updateUserTopic)
	router.Run(":8080")
}
