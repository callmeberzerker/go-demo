package integration

import (
	"github.com/jinzhu/gorm"
)

func GetAllBooks() {

}

func SaveNewBook(entity *Book, db *gorm.DB) *gorm.DB {
	return db.Create(&entity)
}
