package models 

import (
	"book_api/pkg/config"
	"github.com/jinzhu/gorm"

)

var db *gorm.DB

type Book struct{
	gorm.Model
    Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publications string `json:"publications"`
}

func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetByID(ID int64) (*Book , *gorm.DB){
	var getBook Book
	db.Where("ID=?",ID).Find(&getBook)
	return &getBook , db
}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}