package routes

import (
	v1 "user-registration/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.POST("", v1.CreateUser)
		users.GET("/:id", v1.GetUser)
		users.PUT("/:id", v1.UpdateUser)
		users.DELETE("/:id", v1.DeleteUser)
		users.GET("", v1.GetAllUsers)
	}

	return r
}
