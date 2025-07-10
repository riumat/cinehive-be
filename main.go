package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/riumat/cinehive-be/pkg/middleware"
	"github.com/riumat/cinehive-be/pkg/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore nel caricamento del file .env")
	}

	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	app.Listen(":8000")
}
