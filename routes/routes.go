package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(DBMiddleware())
	v1 := r.Group("/api/v1")
	{
		//Test

		//user
		users := v1.Group("/users")

		//restraunt
		restaurants := v1.Group("/restaurant")
		//Test

		//User
		User(users)

		Restaurant(restaurants)

	}
	return r
}
