package main

import (
	"encoding/json"
	"fmt"
	"mindfulness-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTopics(ctx *gin.Context) {
	if database.DB == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}
	rows, err := database.DB.Query("select id, name, description, is_active from topics")
	if err != nil {
		fmt.Println("DB query error:", err)
		ctx.IndentedJSON(400, "Failed to Fetch Contents")

	}
	defer rows.Close()

	topics := []Topic{}
	for rows.Next() {
		var t Topic
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.IsActive)
		if err != nil {
			fmt.Println("Row scan error:", err)
			ctx.IndentedJSON(400, "Failed to Fetch Contents")
			return
		}
		topics = append(topics, t)
	}

	ctx.IndentedJSON(http.StatusOK, topics)
}

func addTopic(ctx *gin.Context) {
	topic := Topic{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Content is not Defined")
		return
	}

	err = json.Unmarshal(data, &topic)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("insert into topics(name, description) values ($1, $2)", topic.Name, topic.Description)
	if err != nil {
		fmt.Println(err)
		ctx.Copy().AbortWithStatusJSON(500, gin.H{"error": "Failed to add topic", "details": err.Error()})
	}
	ctx.IndentedJSON(201, gin.H{"message": "Topic added successfully"})
}
