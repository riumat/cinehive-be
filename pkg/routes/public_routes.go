package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/controllers"
	usercontrollers "github.com/riumat/cinehive-be/app/controllers/userControllers"
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

	route.Get("/movie/:id", controllers.GetMovieDetails)

	route.Get("/tv/:id", controllers.GetTvDetails)

	route.Get("/person/:id", controllers.GetPersonDetails)

	route.Get("/user/movie/:id", middleware.AuthMiddleware(), controllers.GetUserMovieDetails)
	route.Post("/user/movie/:id", middleware.AuthMiddleware(), usercontrollers.AddUserMovie)
	route.Patch("/user/movie/:id", middleware.AuthMiddleware(), usercontrollers.EditUserMovie)
	route.Delete("/user/movie/:id", middleware.AuthMiddleware(), usercontrollers.DeleteUserMovie)

	route.Get("/user/tv/:id", middleware.AuthMiddleware(), controllers.GetUserTvDetails)
	route.Post("/user/tv/:id", middleware.AuthMiddleware(), usercontrollers.AddUserTv)
	route.Patch("/user/tv/:id", middleware.AuthMiddleware(), usercontrollers.EditUserTv)
	route.Delete("/user/tv/:id", middleware.AuthMiddleware(), usercontrollers.DeleteUserTv)

	route.Post("/user/watchlist/movie/:id", middleware.AuthMiddleware(), usercontrollers.AddUserMovieWatchlist)
	route.Delete("/user/watchlist/movie/:id", middleware.AuthMiddleware(), usercontrollers.DeleteUserMovieWatchlist)

	route.Post("/user/watchlist/tv/:id", middleware.AuthMiddleware(), usercontrollers.AddUserTvWatchlist)
	route.Delete("/user/watchlist/tv/:id", middleware.AuthMiddleware(), usercontrollers.DeleteUserTvWatchlist)
}
