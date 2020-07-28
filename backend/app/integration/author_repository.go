package integration

import (
	"github.com/jinzhu/gorm"
)

func createAuthor(entity Author, db *gorm.DB) *gorm.DB {
	return db.Create(&entity)
}
