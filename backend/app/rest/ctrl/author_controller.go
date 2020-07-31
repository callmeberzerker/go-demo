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

// ConfigureAuthorRoutes - configures Author routes
func ConfigureAuthorRoutes(r *mux.Router) {
	r.HandleFunc("/api/author", getAuthors).Methods("GET")
	r.HandleFunc("/api/author/{id}", getAuthor).Methods("GET")
	r.HandleFunc("/api/author", createAuthor).Methods("POST")
	r.HandleFunc("/api/author/{id}", deleteAuthor).Methods("DELETE")
	r.HandleFunc("/api/author/{id}", updateAuthor).Methods("PUT")

	log.Debug().Msg("configured author routes")
}

func getAuthors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	authors, err := service.GetAllAuthors()

	if err != nil {

		log.Error().Msgf("failed creating fetching all authors\nWith reason: [%v]", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(authors)
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	parsedID, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book, err := service.GetAuthorByID(parsedID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(book)
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	savedBook, err := service.CreateBook(book)

	if err != nil {
		log.Error().Msgf("failed creating new book [%#v]\nWith reason: [%v]", book, err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	_ = json.NewEncoder(w).Encode(savedBook)

}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author models.Author
	_ = json.NewDecoder(r.Body).Decode(&author)

	updatedBook, err := service.UpdateAuthor(author)

	if err != nil {

		log.Error().Msgf("failed to update author with id [%d]\nWith reason: [%v]", author.ID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(updatedBook)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.DeleteAuthorByID(uint(ID))

	if err != nil {
		log.Error().Msgf("failed to delete author with id [%d]\nWith reason: [%v]", ID, err)
		w.WriteHeader(http.StatusInternalServerError)

	}

}
