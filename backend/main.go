package main

import (
	"github.com/gorilla/mux"
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/ctrl"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	db := integration.OpenDbConnection()

	db.AutoMigrate(&integration.Book{})
	db.AutoMigrate(&integration.Author{})
	db.Model(&integration.Book{}).AddForeignKey("author_id", "author(id)", "RESTRICT", "RESTRICT")

	ctrl.Configure(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
