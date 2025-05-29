package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/riumat/cinehive-be/migrations"
	"github.com/riumat/cinehive-be/pkg/middleware"
	"github.com/riumat/cinehive-be/pkg/routes"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrations.Migrate()
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore nel caricamento del file .env")
	}

	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	app.Listen(":8000")
}
