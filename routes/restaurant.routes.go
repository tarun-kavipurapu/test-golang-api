package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tarun-kavipurapu/gin-gorm/controllers"
	"github.com/tarun-kavipurapu/gin-gorm/middleware"
)

func Restaurant(r *gin.RouterGroup) {
	r.POST("/createRestaurant", middleware.JWTRestaurantAuth(), controllers.CreateRestaurant)
}
