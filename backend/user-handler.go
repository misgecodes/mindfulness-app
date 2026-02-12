package main

import (
	"encoding/json"
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

func addUser(c *gin.Context) {
	body := User{}
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "User is not Defined")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	_, err = database.DB.Exec("insert into users(id, username,password) values ($1, $2, $3)", body.ID, body.Username, body.Password)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create user", "details": err.Error()})
	} else {
		fmt.Println("User created successfully")
		c.IndentedJSON(201, gin.H{"message": "User created successfully"})
	}

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
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			fmt.Println("Row scan error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading user data"})
			return
		}
		users = append(users, u)
	}

	ctx.JSON(http.StatusOK, users)
}
