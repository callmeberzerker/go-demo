package integration

import (
	"github.com/jinzhu/gorm"
)

// GetAllBooks - gets all books
func GetAllBooks() {

}

// SaveNewBook - saves a new book
func SaveNewBook(entity *Book, db *gorm.DB) *gorm.DB {
	return db.Create(&entity)
}
