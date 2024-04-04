// user.routes.go

package routes

import (
	// Import your controller package

	"github.com/gin-gonic/gin"
	"github.com/tarun-kavipurapu/gin-gorm/controllers"
)

func User(r *gin.RouterGroup) {
	// Define routes for user-related endpoints
	r.GET("/test", controllers.Test)

	r.GET("/user", controllers.FindUsers) // Handler for fetching users
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LoginUser)
	// Handler for creating users
	// r.GET("/roles", controllers.SeedRole)
}
