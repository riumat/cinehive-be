package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func GetMovieDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	client := config.NewTMDBClient()
	data, err := services.FetchMovieDetails(client, id)
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

func GetMovieCast(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	client := config.NewTMDBClient()
	data, err := services.FetchMovieCastDetails(client, id)
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

func GetMovieCrew(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	client := config.NewTMDBClient()
	data, err := services.FetchMovieCrewDetails(client, id)
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

func GetMovieVideos(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	client := config.NewTMDBClient()
	data, err := services.FetchMovieVideos(client, id)
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

func GetMovieRecommendations(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	client := config.NewTMDBClient()
	data, err := services.FetchMovieRecommendations(client, id)
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
