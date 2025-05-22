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

	route.Get("/movie/featured", controllers.GetFeaturedMovie)

	route.Get("/movie/:id/details", controllers.GetMovieDetails)
	route.Get("/movie/:id/cast", controllers.GetMovieCast)
	route.Get("/movie/:id/crew", controllers.GetMovieCrew)
	route.Get("/movie/:id/videos", controllers.GetMovieVideos)
	route.Get("/movie/:id/recommendations", controllers.GetMovieRecommendations)

	route.Get("/tv/:id/details", controllers.GetTvDetails)
	route.Get("/tv/:id/cast", controllers.GetTvCast)
	route.Get("/tv/:id/crew", controllers.GetTvCrew)
	route.Get("/tv/:id/videos", controllers.GetTvVideos)
	route.Get("/tv/:id/recommendations", controllers.GetTvRecommendations)
	route.Get("/tv/:id/seasons", controllers.GetTvSeasons)

	route.Get("/person/:id", controllers.GetPersonDetails)
}
