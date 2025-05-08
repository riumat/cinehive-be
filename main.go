package main

import (
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/database"
	"github.com/riumat/cinehive-be/migrations"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrations.Migrate()
		return
	}

	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Cinehive API")
	})

	app.Listen(":8000")
	app.Use(swagger.New())
}
