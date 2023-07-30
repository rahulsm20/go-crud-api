package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("X-Serverless-Function", "GinServer")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Access-Control-Allow-Origin", "*")

	engine := gin.New()

	// Recover from any panics
	engine.Use(gin.Recovery())

	routes.PostRoutes(engine.Group("/api"))
	routes.UserRoutes(engine.Group("/api"))

	engine.ServeHTTP(w, r)
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
