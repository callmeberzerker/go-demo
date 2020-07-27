package integration

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Isbn     string `gorm:"type:varchar(100);unique_index"`
	Title    string `gorm:"size:255"`
	Author   Author `gorm:"foreignkey:AuthorId"`
	AuthorId uint
}

type Author struct {
	gorm.Model
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
}

func (Book) TableName() string {
	return "book"
}

func (Author) TableName() string {
	return "author"
}
