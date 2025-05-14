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

	route.Get("/movie/featured", controllers.GetFeaturedMovie)

	route.Get("/movie/:id/header", controllers.GetMovieHeader)
	route.Get("/movie/:id/overview", controllers.GetMovieOverview)
	route.Get("/movie/:id/cast", controllers.GetMovieCast)
	route.Get("/movie/:id/crew", controllers.GetMovieCrew)
	route.Get("/movie/:id/videos", controllers.GetMovieVideos)
	route.Get("/movie/:id/recommendations", controllers.GetMovieRecommendations)
}
