package main

import (
	"github.com/rahulsm20/go-crud-api/initializers"
	"github.com/rahulsm20/go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
