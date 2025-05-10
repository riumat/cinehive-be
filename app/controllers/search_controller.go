package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
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
