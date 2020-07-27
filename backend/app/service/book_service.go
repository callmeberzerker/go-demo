package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/models"
)

func SaveBook(bookDto models.Book) models.Book {
	db := integration.OpenDbConnection()

	trx := db.Begin()

	var bookEntity = &integration.Book{
		Isbn:  bookDto.Isbn,
		Title: bookDto.Title,
		Author: integration.Author{
			Model:     gorm.Model{ID: bookDto.Author.ID},
			FirstName: bookDto.Author.FirstName,
			LastName:  bookDto.Author.LastName,
		},
	}
	if result := integration.SaveNewBook(bookEntity, trx); result.Error != nil {

		fmt.Println("We have error!")
		fmt.Println(result.Error)
		trx.Rollback()
	}

	trx.Commit()
	db.Close()
	return models.Book{}
}
