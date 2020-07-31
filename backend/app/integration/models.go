package integration

import (
	"github.com/jinzhu/gorm"
)

// Book - Entity
type Book struct {
	gorm.Model
	Isbn     string `gorm:"type:varchar(100);unique_index"`
	Title    string `gorm:"size:255"`
	Author   Author `gorm:"foreignkey:AuthorID"`
	AuthorID uint
}

// Author - Entity
type Author struct {
	gorm.Model
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
}

// TableName - ensures proper SQL table name
func (Book) TableName() string {
	return "books"
}

// TableName - ensures proper SQL table name
func (Author) TableName() string {
	return "authors"
}
