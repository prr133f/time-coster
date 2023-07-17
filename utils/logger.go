package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func InitLogger() {
	switch strings.ToLower(os.Getenv("APP_LEVEL")) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "prod":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
