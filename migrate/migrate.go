package main

import (
	"github.com/fahad-md-kamal/go-crud/initializers"
	"github.com/fahad-md-kamal/go-crud/models"
)


func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}