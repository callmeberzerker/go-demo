package ctrl

import (
	"encoding/json"
	"go-demo/main/app/rest/models"
	"go-demo/main/app/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var books []models.Book

// ConfigureBookRoutes - configures book routes
func ConfigureBookRoutes(r *mux.Router) {
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")

	log.Debug().Msg("configured book routes")
}

func getBooks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	log.Info().Msg("Info message")
	books, err := service.GetAllBooks()

	if err != nil {
		log.Printf("ERROR: Failed fetching books %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	parsedID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := service.GetBook(parsedID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	savedBook, err := service.CreateBook(book)

	if err != nil {
		log.Error().Msgf("failed creating new book [%#v]", book)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	_ = json.NewEncoder(w).Encode(savedBook)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	updatedBook, err := service.UpdateBook(book)

	if err != nil {

		log.Printf("Failed to update book with id [%d]\n", book.ID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(updatedBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)

	// for i, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:i], books[i+1:]...)
	// 		break
	// 	}
	// }

	// _ = json.NewEncoder(w).Encode(books)

}
