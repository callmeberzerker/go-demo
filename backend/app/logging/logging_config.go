package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ConfigLogging - config logging
func ConfigLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
