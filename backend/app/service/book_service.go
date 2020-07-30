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
	bookEntity := &integration.Book{
		Isbn:     bookDto.Isbn,
		Title:    bookDto.Title,
		AuthorID: uint(bookDto.AuthorID),
	}
	err := db.Transaction(func(trx *gorm.DB) error {

		if err := integration.SaveNewBook(bookEntity, trx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Book{}, err
	}

	return models.Book{
		ID:       int64(bookEntity.ID),
		AuthorID: int64(bookEntity.AuthorID),
		Title:    bookEntity.Title,
		Isbn:     bookEntity.Isbn,
	}, err
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
			ID:       int64(book.ID),
			Isbn:     book.Isbn,
			Title:    book.Title,
			AuthorID: int64(book.AuthorID),
		}
		booksMapped = append(booksMapped, bookRest)

	}

	return booksMapped, nil
}

// GetBook - Gets a single book by id
func GetBook(ID int64) (models.Book, error) {

	db := integration.OpenDbConnection()

	book, err := integration.GetBookByID(ID, db)
	if err != nil {
		return models.Book{}, db.Error
	}

	bookRest := models.Book{
		ID:       int64(book.ID),
		AuthorID: int64(book.AuthorID),
		Isbn:     book.Isbn,
		Title:    book.Title,
	}

	return bookRest, nil

}

// UpdateBook - updates a book
func UpdateBook(bookDto models.Book) (models.Book, error) {

	db := integration.OpenDbConnection()
	bookEntity := &integration.Book{
		Model:    gorm.Model{ID: uint(bookDto.ID)},
		Isbn:     bookDto.Isbn,
		Title:    bookDto.Title,
		AuthorID: uint(bookDto.AuthorID),
	}
	err := db.Transaction(func(trx *gorm.DB) error {

		if err := integration.UpdateBook(bookEntity, trx); err != nil {

			return err
		}

		return nil
	})

	if err != nil {
		return models.Book{}, err
	}

	updatedBook := models.Book{
		ID:       int64(bookEntity.ID),
		Isbn:     bookEntity.Isbn,
		Title:    bookEntity.Title,
		AuthorID: int64(bookEntity.AuthorID),
	}

	return updatedBook, nil

}

// DeleteBook - deletes a book by id
func DeleteBook(ID uint) error {
	db := integration.OpenDbConnection()

	err := integration.DeleteBook(ID, db)

	return err
}
