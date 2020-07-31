package service

import (
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/models"

	"github.com/jinzhu/gorm"
)

// CreateAuthor - creates a new author
func CreateAuthor(authorDto models.Author) (models.Author, error) {
	db := integration.OpenDbConnection()
	author := &integration.Author{
		FirstName: authorDto.FirstName,
		LastName:  authorDto.LastName,
	}
	err := db.Transaction(func(trx *gorm.DB) error {

		if err := integration.SaveNewAuthor(author, trx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Author{}, err
	}

	return models.Author{
		ID:        int64(author.ID),
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}, err
}

// GetAllAuthors - gets all authors
func GetAllAuthors() ([]models.Author, error) {

	db := integration.OpenDbConnection()
	authorsFromDB, err := integration.GetAllAuthors(db)
	var authorsMapped = []models.Author{}

	if err != nil {
		return []models.Author{}, err
	}

	for _, author := range authorsFromDB {

		authorRest := models.Author{
			ID:        int64(author.ID),
			FirstName: author.FirstName,
			LastName:  author.LastName,
		}
		authorsMapped = append(authorsMapped, authorRest)

	}

	return authorsMapped, nil
}

// GetAuthorByID - Gets a single author by id
func GetAuthorByID(ID int64) (models.Author, error) {

	db := integration.OpenDbConnection()

	author, err := integration.GetAuthorByID(ID, db)
	if err != nil {
		return models.Author{}, db.Error
	}

	authorRest := models.Author{
		ID:        int64(author.ID),
		FirstName: author.FirstName,
		LastName:  author.LastName,
	}

	return authorRest, nil

}

// UpdateAuthor - updates author
func UpdateAuthor(authorDto models.Author) (models.Author, error) {

	db := integration.OpenDbConnection()
	authorEntity := &integration.Author{
		Model:     gorm.Model{ID: uint(authorDto.ID)},
		FirstName: authorDto.FirstName,
		LastName:  authorDto.LastName,
	}
	err := db.Transaction(func(trx *gorm.DB) error {

		if err := integration.UpdateAuthor(authorEntity, trx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return models.Author{}, err
	}

	updatedAuthor := models.Author{
		ID:        int64(authorEntity.ID),
		FirstName: authorEntity.FirstName,
		LastName:  authorEntity.LastName,
	}

	return updatedAuthor, nil

}

// DeleteAuthorByID - deletes author by ID
func DeleteAuthorByID(ID uint) error {
	db := integration.OpenDbConnection()

	err := integration.DeleteAuthorByID(ID, db)

	return err
}
