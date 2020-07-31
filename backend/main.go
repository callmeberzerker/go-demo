package main

import (
	"go-demo/main/app/integration"
	"go-demo/main/app/logging"
	"go-demo/main/app/rest/ctrl"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	logging.ConfigLogging()
	r := mux.NewRouter()
	integration.MigrateModels()
	ctrl.ConfigureBookRoutes(r)
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	log.Error().Err(http.ListenAndServe(":8080", handlers.CORS(allowedMethods, allowedHeaders)(r)))
}
