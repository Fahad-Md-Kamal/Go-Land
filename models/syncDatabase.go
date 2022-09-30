package models

import "github.com/fahad-md-kamal/go-jwt/initializers"


func SyncDatabase() {
	initializers.DB.AutoMigrate(&User{})
}