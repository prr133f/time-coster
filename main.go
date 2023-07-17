package main

import (
	"os"
	"time-coster/routes"
	"time-coster/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	routes.InitRoutes(app)

	utils.InitLogger()

	log.Fatal().Err(app.Listen(os.Getenv("APP_PORT")))
}
