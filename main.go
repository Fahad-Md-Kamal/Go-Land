package main

import (
	"github.com/fahad-md-kamal/go-jwt/controllers"
	"github.com/fahad-md-kamal/go-jwt/initializers"
	"github.com/fahad-md-kamal/go-jwt/middlewares"
	"github.com/fahad-md-kamal/go-jwt/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	models.SyncDatabase()
}


func main() {
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run()
}