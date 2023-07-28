package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/controllers"
	"github.com/rahulsm20/go-crud-api/pkg/middleware"
)

func PostRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.GET("/:id", middleware.RequireAuth, controllers.FetchPostByID)
		posts.PUT("/:id", middleware.RequireAuth, controllers.UpdatePost)
		posts.DELETE("/:id", middleware.RequireAuth, controllers.DeletePost)
		posts.POST("/", middleware.RequireAuth, controllers.CreatePost)
		posts.GET("/", middleware.RequireAuth, controllers.FetchAllPosts)
	}
}
