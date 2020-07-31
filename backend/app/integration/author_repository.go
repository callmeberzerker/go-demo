package integration

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// GetAllAuthors - gets all authors
func GetAllAuthors(db *gorm.DB) ([]Author, error) {

	var authors []Author
	dbResult := db.Find(&authors)
	return authors, dbResult.Error

}

// UpdateAuthor - updates author
func UpdateAuthor(entity *Author, trx *gorm.DB) error {
	isNewRecord := trx.NewRecord(entity) // => returns `true` as primary key is blank
	var error error
	if !isNewRecord {
		error = trx.Model(&entity).Update(&entity).Error
	} else {
		error = errors.New("Can't update a record with non-existent primary key")
	}

	return error
}

// SaveNewAuthor - saves a new author
func SaveNewAuthor(entity *Author, db *gorm.DB) error {
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

// GetAuthorByID - gets author by id
func GetAuthorByID(id int64, db *gorm.DB) (Author, error) {
	author := Author{}
	var err error

	if db.First(&author, id).RecordNotFound() && db.Error != nil {
		err = errors.New("Record not found")
	}

	return author, err
}

// DeleteAuthorByID - updates a book
func DeleteAuthorByID(ID uint, db *gorm.DB) error {
	if err := db.Delete(&Book{Model: gorm.Model{ID: ID}}).Error; err != nil {
		return err
	}
	log.Info().Msgf("Deleted book with id [%d]", ID)
	return nil
}
