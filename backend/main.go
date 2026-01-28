package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10/translations/id"
)

func main() {
	router := gin.Default()
	router.GET("/topics", getTopics)
	router.Run("localhost:8080")
}
