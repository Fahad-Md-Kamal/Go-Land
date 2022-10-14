package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Goly struct{
	ID			uint64 	`json:"id" gorm:"primaryKey"`
	Redirect 	string 	`json:"redirect"`
	Goly 		string	`json:"goly" gorm:"unique;not null"`
	Clicked		uint64 	`json:"clicked"`
	Random 		bool 	`json:"random"`
}

func Setup() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=fiber_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	err = DB.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}