package models

import (
	"fahadmdkamal.com/booksapi/pkg/config"
	"gorm.io/gorm"
)

var db * gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}


func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}


func GetBookById(Id int64) (*Book, *gorm.DB ){
	var Book Book
	db.Where("ID=?", Id).Find(&Book)
	return &Book, db
}


func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}