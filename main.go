package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/routes"
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

	routes.PostRoutes(r.Group(""))
	routes.UserRoutes(r.Group(""))
	r.Run()
}

// func main() {
// 	lambda.Start(Handler)
// }
