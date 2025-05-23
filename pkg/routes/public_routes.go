package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/controllers"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/auth/signup", controllers.UserSignUp)
	route.Post("/auth/signin", controllers.UserSignIn)

	route.Get("/trending", controllers.GetTrendingContent)

	route.Get("/genres/movie", controllers.GetMovieGenres)
	route.Get("/genres/tv", controllers.GetTvGenres)

	route.Get("/search", controllers.GetSearchResults)
	route.Get("/search/filters", controllers.GetSearchWithFilters)

	route.Get("/movie/:id/details", controllers.GetMovieDetails)

	route.Get("/tv/:id/details", controllers.GetTvDetails)

	route.Get("/person/:id", controllers.GetPersonDetails)
}
