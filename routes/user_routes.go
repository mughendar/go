package routes

import (
	"database/sql"
	"my-gin-mysql-app/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up the routes for user-related actions
func SetupUserRoutes(router *gin.Engine, db *sql.DB) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", controllers.CreateUser(db))
	}
}
