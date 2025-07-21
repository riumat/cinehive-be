package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/riumat/cinehive-be/pkg/middleware"
	"github.com/riumat/cinehive-be/pkg/routes"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println("cant load .env file :", err)
		} else {
			log.Println(".env file loaded successfully")
		}
	} else {
		log.Println("No .env file found, using environment variables")
	}

	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	app.Listen(":8000")
}
