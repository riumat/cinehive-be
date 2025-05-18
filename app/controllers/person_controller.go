package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func GetPersonDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	client := config.NewTMDBClient()
	data, err := services.FetchPersonDetails(client, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(data)
}
