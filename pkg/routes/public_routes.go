package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/controllers"
	"github.com/riumat/cinehive-be/pkg/middleware"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/auth/signup", controllers.SupabaseUserSignUp)
	route.Post("/auth/signin", controllers.SupabaseUserSignIn)
	route.Post("/auth/refresh-token", controllers.RefreshAuthToken)

	route.Get("/trending", controllers.GetTrendingContent)

	route.Get("/genres/movie", controllers.GetMovieGenres)
	route.Get("/genres/tv", controllers.GetTvGenres)

	route.Get("/search", controllers.GetSearchResults)
	route.Get("/search/filters", controllers.GetSearchWithFilters)

	route.Get("/movie/:id/details", middleware.AuthMiddleware(), controllers.GetMovieDetails)

	route.Get("/tv/:id/details", middleware.AuthMiddleware(), controllers.GetTvDetails)

	route.Get("/person/:id", controllers.GetPersonDetails)

	route.Post("/user/content", middleware.AuthMiddleware(), controllers.AddUserContent)
}
