package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	//route.Get("/books", controllers.GetBooks)
	//route.Get("/book/:id", controllers.GetBook)

	route.Post("/auth/signup", controllers.UserSignUp)
	route.Post("/auth/signin", controllers.UserSignIn)
}
