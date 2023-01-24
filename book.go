package models

import (
	"github.com/balaram/bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`             //structure
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	config.Connect()                    //this is a function
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
                                              //this all are batabase connection part 
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {

	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {

	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {

	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
