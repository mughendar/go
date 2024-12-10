package controllers

import (
	"database/sql"
	"my-gin-mysql-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user
func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		// Bind the JSON payload to the user struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the user into the database
		query := "INSERT INTO users (name, email) VALUES (?, ?)"
		result, err := db.Exec(query, user.Name, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
			return
		}

		// Get the last inserted ID
		id, _ := result.LastInsertId()
		user.ID = int(id)

		// Respond with the created user
		c.JSON(http.StatusOK, gin.H{
			"message": "User successfully added!",
			"user":    user,
		})
	}
}
