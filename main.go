package main

import (
	"github.com/fahad-md-kamal/go-jwt/controllers"
	"github.com/fahad-md-kamal/go-jwt/initializers"
	"github.com/fahad-md-kamal/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)


func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}


func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup )
	r.POST("/login", controllers.Login )
	r.GET("/validate", middleware.RequrieAuth, controllers.Validate )
	r.Run()
}