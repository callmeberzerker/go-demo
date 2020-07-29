package main

import (
	"go-demo/main/app/integration"
	"go-demo/main/app/rest/ctrl"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	integration.MigrateModels()

	ctrl.ConfigureBookRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
