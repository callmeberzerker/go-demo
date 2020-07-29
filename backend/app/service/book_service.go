package service

import (
	"fmt"
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/models"

	"github.com/jinzhu/gorm"
)

// CreateBook - creates a new book
func CreateBook(bookDto models.Book) (models.Book, error) {
	db := integration.OpenDbConnection()

	err := db.Transaction(func(trx *gorm.DB) error {
		// var bookEntity = &integration.Book{
		// 	Isbn:  bookDto.Isbn,
		// 	Title: bookDto.Title,
		// 	Author: integration.Author{
		// 		Model:     gorm.Model{ID: bookDto.Author.ID},
		// 		FirstName: bookDto.Author.FirstName,
		// 		LastName:  bookDto.Author.LastName,
		// 	},
		// }
		// if result, err := integration.SaveNewBook(bookEntity, trx); err != nil || result.Error != nil {
		// 	return err
		// }

		return nil
	})

	if err != nil {
		return models.Book{}, err
	}

	return models.Book{}, nil
}

// GetAllBooks - gets all books
func GetAllBooks() ([]models.Book, error) {

	db := integration.OpenDbConnection()
	booksFromDb, err := integration.GetAllBooks(db)
	var booksMapped = []models.Book{}

	if err != nil {
		fmt.Printf("Error happened %v", err)
		return []models.Book{}, err
	}

	for _, book := range booksFromDb {

		bookRest := models.Book{
			ID:    int64(book.ID),
			Isbn:  book.Isbn,
			Title: book.Title,
		}
		booksMapped = append(booksMapped, bookRest)

	}

	fmt.Println("returning")
	return booksMapped, nil
}

// GetBook - Gets a single book by id
func GetBook(ID int64) (*models.Book, error) {

	db := integration.OpenDbConnection()

	book := integration.GetBookByID(ID, db)
	if db.Error != nil || book == nil {
		return nil, db.Error
	}

	bookRest := models.Book{
		ID: int64(book.ID),
	}

	return &bookRest, nil

}
