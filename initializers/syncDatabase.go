package initializers

import "github.com/fahad-md-kamal/go-jwt/models"


func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}