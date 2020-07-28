package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/models"
)

func CreateBook(bookDto models.Book) (models.Book, error) {
	db := integration.OpenDbConnection()

	err := db.Transaction(func(trx *gorm.DB) error {
		var bookEntity = &integration.Book{
			Isbn:  bookDto.Isbn,
			Title: bookDto.Title,
			Author: integration.Author{
				Model:     gorm.Model{ID: bookDto.Author.ID},
				FirstName: bookDto.Author.FirstName,
				LastName:  bookDto.Author.LastName,
			},
		}
		if err := integration.SaveNewBook(bookEntity, trx).Error; err != nil {

			return err
		}
		return nil
	})

	if err != nil {
		return models.Book{}, err
	}

	return models.Book{}, nil
}
