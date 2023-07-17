package routes

import (
	"context"
	"os"
	"time-coster/domain"
	v1 "time-coster/routes/v1"
	"time-coster/view"

	"github.com/gofiber/fiber/v2"
	log "github.com/rs/zerolog/log"
)

func InitRoutes(app *fiber.App) {
	pg, err := domain.NewPG(context.Background(), os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal().Err(err)
	}
	router := v1.Router{App: app, View: &view.View{Pg: pg}}
	router.Router()
}
