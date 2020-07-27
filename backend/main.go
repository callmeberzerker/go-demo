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

	integration.MigrateModels()

	ctrl.ConfigureBookRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
