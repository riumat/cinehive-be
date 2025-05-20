package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func GetFeaturedMovie(c *fiber.Ctx) error {
	client := config.NewTMDBClient()
	data, err := services.FetchFeaturedMovie(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  data,
	})
}

func GetTrendingContent(c *fiber.Ctx) error {
	client := config.NewTMDBClient()
	data, err := services.FetchLandingCards(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"trendingMovies": data["trendingMovies"],
			"trendingTv":     data["trendingTv"],
			"topRatedMovies": data["topRatedMovies"],
			"topRatedTv":     data["topRatedTv"],
		},
	})
}
