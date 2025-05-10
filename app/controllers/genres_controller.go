package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const movieType = "movie"
const tvType = "tv"

func GetMovieGenres(c *fiber.Ctx) error {
	client := config.NewTMDBClient()

	data, err := services.FetchGenres(client, movieType)
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

func GetTvGenres(c *fiber.Ctx) error {
	client := config.NewTMDBClient()

	data, err := services.FetchGenres(client, tvType)
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
