package main

import (
	"go-demo/main/app/integration"
	"go-demo/main/app/logging"
	"go-demo/main/app/rest/ctrl"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logging.ConfigLogging()
	r := mux.NewRouter()
	integration.MigrateModels()
	ctrl.ConfigureBookRoutes(r)
	log.Debug().Msg("Heyo")
	log.Error().Err(http.ListenAndServe(":8080", handlers.CORS()(r)))
}
