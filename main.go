package main

import (
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/migrations"
	"github.com/riumat/cinehive-be/pkg/middleware"
	"github.com/riumat/cinehive-be/pkg/routes"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrations.Migrate()
		return
	}

	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	app.Listen(":8000")
	app.Use(swagger.New())
}
