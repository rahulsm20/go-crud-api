package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/controllers"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	posts := r.Group("/posts")
	{
		posts.GET("/:id", controllers.FetchPostByID)
		posts.PUT("/:id", controllers.UpdatePost)
		posts.DELETE("/:id", controllers.DeletePost)
		posts.POST("/", controllers.CreatePost)
		posts.GET("/", controllers.FetchAllPosts)
	}

	users := r.Group("/users")
	{
		users.POST("/", controllers.Signup)
		users.DELETE("/", controllers.DeleteUser)
	}
	r.Run()
}
