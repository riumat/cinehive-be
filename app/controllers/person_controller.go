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
	return c.JSON(fiber.Map{
		"data":  data,
		"error": false,
	})
}

func GetUserPersonDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	token := c.Locals("token")
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

	user := c.Locals("user_id").(string)
	supabaseClient := config.NewSupabaseClient(token.(string))
	data, err := services.FetchPersonUserData(supabaseClient, user, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch user movie details: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  data,
		"error": false,
	})
}
