package integration

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// GetAllBooks - gets all books
func GetAllBooks(db *gorm.DB) ([]Book, error) {

	var books []Book
	dbResult := db.Find(&books)
	return books, dbResult.Error

}

// UpdateBook - updates a book
func UpdateBook(entity *Book, trx *gorm.DB) error {
	isNewRecord := trx.NewRecord(entity) // => returns `true` as primary key is blank
	var error error
	if !isNewRecord {
		error = trx.Model(&entity).Update(&entity).Error
	} else {
		error = errors.New("Can't update a record with non-existent primary key")
	}

	return error
}

// SaveNewBook - saves a new book
func SaveNewBook(entity *Book, db *gorm.DB) error {
	isNewRecord := db.NewRecord(entity) // => returns `true` as primary key is blank
	var err error

	if isNewRecord {
		if dbErr := db.Create(&entity).Error; dbErr != nil {
			err = dbErr
		}
	} else {
		err = errors.New("cannot create a new entity with a primary key set")
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

// DeleteBook - updates a book
func DeleteBook(ID uint, db *gorm.DB) error {
	if err := db.Delete(&Book{Model: gorm.Model{ID: ID}}).Error; err != nil {
		return err
	}
	log.Info().Msgf("Deleted book with id [%d]", ID)
	return nil
}
