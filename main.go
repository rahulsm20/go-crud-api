package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/controllers"
	"github.com/rahulsm20/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/:id", controllers.FetchPostByID)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.FetchAllPosts)
	r.Run()
}
