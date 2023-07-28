package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("/login", controllers.Signin)
		users.POST("/", controllers.Signup)
		users.DELETE("/", controllers.DeleteUser)
	}
}
