package models

//importing gorm and config file
import (
	"github.com/MostafaAbdelkarim/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//building the Book entity using struct and gorm.model
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//implementing init function
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Implementing reciever function CreateBook of type Book that returns book by reference.
func (b *Book) CreateBook() *Book {
	//NewRecord inner functionality in GORM allowing us to skip boilerplate sql queries
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//Implementing GetAllBooks method returning a slice named books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//Implementing GetBookById function which returns both Book reference from Db and DB from gorm
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID?=", Id).Find(&getBook)
	return &getBook, db
}

//Implementing the DeleteBook method
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
