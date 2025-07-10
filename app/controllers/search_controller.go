package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

func GetSearchResults(c *fiber.Ctx) error {
	name := c.Query("query")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Query parameter is required",
		})
	}

	page := c.Query("page", "1")

	client := config.NewTMDBClient()

	data, err := services.FetchSearchResults(client, name, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"data":          data.Results,
		"page":          data.Page,
		"total_pages":   data.TotalPages,
		"total_results": data.TotalResults,
	})
}

func GetMovieSearchResults(c *fiber.Ctx) error {
	name := c.Query("query")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Query parameter is required",
		})
	}

	page := c.Query("page", "1")

	client := config.NewTMDBClient()

	data, err := services.FetchMovieSearchResults(client, name, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"data":          data.Results,
		"page":          data.Page,
		"total_pages":   data.TotalPages,
		"total_results": data.TotalResults,
	})
}

func GetTvSearchResults(c *fiber.Ctx) error {
	name := c.Query("query")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Query parameter is required",
		})
	}

	page := c.Query("page", "1")

	client := config.NewTMDBClient()

	data, err := services.FetchTvSearchResults(client, name, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"data":          data.Results,
		"page":          data.Page,
		"total_pages":   data.TotalPages,
		"total_results": data.TotalResults,
	})
}

func GetPersonSearchResults(c *fiber.Ctx) error {
	name := c.Query("query")
	if name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Query parameter is required",
		})
	}

	page := c.Query("page", "1")

	client := config.NewTMDBClient()

	data, err := services.FetchPersonSearchResults(client, name, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"data":          data.Results,
		"page":          data.Page,
		"total_pages":   data.TotalPages,
		"total_results": data.TotalResults,
	})
}

func GetSearchWithFilters(c *fiber.Ctx) error {
	media := c.Query("media")
	if media == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Media parameter (movie or tv) is required",
		})
	}

	currentYear := time.Now().Format("2006")

	params := types.FilterParams{
		Genres:     c.Query("genres", ""),
		Providers:  c.Query("providers", ""),
		Page:       c.Query("page", "1"),
		From:       c.Query("from", "1900"),
		To:         c.Query("to", currentYear),
		Sort:       c.Query("sort", "popularity.desc"),
		RuntimeGte: c.Query("runtime_gte", "0"),
		RuntimeLte: c.Query("runtime_lte", "400"),
	}

	client := config.NewTMDBClient()

	data, err := services.FetchSearchWithFilters(client, params, media)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"data":          data.Results,
		"page":          data.Page,
		"total_pages":   data.TotalPages,
		"total_results": data.TotalResults,
	})
}
