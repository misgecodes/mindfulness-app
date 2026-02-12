package main

import (
	"encoding/json"
	// "fmt"
	"fmt"
	"mindfulness-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getContents(ctx *gin.Context) {
	id := ctx.Param("topic-id")
	if database.DB == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}
	rows, err := database.DB.Query("select id, topic_id, title, uri from contents where topic_id = $1", id)

	if err != nil {
		fmt.Println("DB query error:", err)
		ctx.IndentedJSON(400, "Failed to Fetch Contents")
		return
	}

	defer rows.Close()

	contents := []Content{}
	for rows.Next() {
		var c Content
		err := rows.Scan(&c.ID, &c.TopicID, &c.Title, &c.URI)
		if err != nil {
			fmt.Println("Row scan error:", err)
			ctx.IndentedJSON(400, "Failed to Fetch Contents")
			return
		}
		contents = append(contents, c)
	}
	if len(contents) == 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "No Contents Found for the Given Topic ID"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, contents)
}

func addContent(ctx *gin.Context) {
	content := Content{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Content is not Defined")
		return
	}
	err = json.Unmarshal(data, &content)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("insert into contents(topic_id, title, uri) values ($1, $2, $3)", content.TopicID, content.Title, content.URI)
	if err != nil {
		fmt.Println(err)
		ctx.Copy().AbortWithStatusJSON(500, gin.H{"error": "Failed to add content", "details": err.Error()})
	} else {
		ctx.IndentedJSON(201, gin.H{"message": "Content added successfully"})
	}
}
