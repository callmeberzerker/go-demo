package integration

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// GetAllBooks - gets all books
func GetAllBooks(db *gorm.DB) ([]Book, error) {

	var books []Book
	dbResult := db.Find(&books)
	return books, dbResult.Error

}

// UpdateBook - updates a book
func UpdateBook(entity *Book, db *gorm.DB) (*gorm.DB, error) {
	isNewRecord := db.NewRecord(entity) // => returns `true` as primary key is blank

	if !isNewRecord {
		return db.Create(&entity), nil
	}

	return db, errors.New("Can't create a record with non-existent primary key")
}

// SaveNewBook - saves a new book
func SaveNewBook(entity *Book, db *gorm.DB) error {
	isNewRecord := db.NewRecord(entity) // => returns `true` as primary key is blank
	var err error

	if isNewRecord {
		db.Create(&entity)
	} else {
		err = errors.New("cannot create a new entity with a primary key set")
	}

	if db.Error != nil {
		err = db.Error
	}

	return err
}

// GetBookByID - gets a book by id
func GetBookByID(id int64, db *gorm.DB) (Book, error) {
	book := Book{}
	var err error

	if db.First(&book, id).RecordNotFound() && db.Error != nil {
		err = errors.New("Record not found")
	}

	return book, err
}
