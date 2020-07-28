package ctrl

import (
	"encoding/json"
	"fmt"
	"go-demo/main/app/rest/models"
	"go-demo/main/app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var books []models.Book

// ConfigureBookRoutes - configures book routes
func ConfigureBookRoutes(r *mux.Router) {
	books = append(books, models.Book{ID: "1", Isbn: "1111111", Title: "Book1", Author: &models.Author{
		FirstName: "John",
		LastName:  "Doe",
	}})
	books = append(books, models.Book{ID: "2", Isbn: "3333", Title: "Book2", Author: &models.Author{
		FirstName: "Here we have stuff",
		LastName:  "Doe",
	}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
}

func getBooks(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("333")
	fmt.Println("2 11")
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	_ = json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	savedBook, err := service.CreateBook(book)

	if err != nil {
		log.Printf("ERROR: Failed saving %v", book)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	_ = json.NewEncoder(w).Encode(savedBook)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	_ = json.NewEncoder(w).Encode(books)

}
