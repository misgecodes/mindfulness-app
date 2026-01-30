package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, topics)
}

func createUserTopic(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User topic created"})
}
