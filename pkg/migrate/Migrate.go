package main

import (
	"github.com/rahulsm20/go-crud-api/pkg/initializers"
	"github.com/rahulsm20/go-crud-api/pkg/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
}
