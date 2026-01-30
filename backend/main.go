package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10/translations/id"
)

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/topics", getTopics)
	router.GET("/health", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"status": "OK"})
	})
	router.Run(":8080")
}
