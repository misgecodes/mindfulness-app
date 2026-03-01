package main

import (
	// "encoding/json"
	"fmt"
	"mindfulness-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUserTopic(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User topic created"})
}

func updateUserTopic(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User topic updated"})
}

func getUsers(ctx *gin.Context) {
	if database.DB == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	rows, err := database.DB.Query("SELECT id, username, password FROM users")
	if err != nil {
		fmt.Println("DB query error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Email, &u.Password)
		if err != nil {
			fmt.Println("Row scan error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading user data"})
			return
		}
		users = append(users, u)
	}

	ctx.JSON(http.StatusOK, users)
}
