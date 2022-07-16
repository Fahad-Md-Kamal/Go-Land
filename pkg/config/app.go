package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

  var (
	db * gorm.DB
  )

  func Connect() {
	dsn := "host=localhost user=genex password=postgres dbname=book_shop_db port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db = d
  }

func GetDB() *gorm.DB{
	return db
}