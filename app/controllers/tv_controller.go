package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func GetTvDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Movie ID is required",
		})
	}
	tmdbClient := config.NewTMDBClient()
	data, err := services.FetchTvDetails(tmdbClient, id)
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

func GetUserTvDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "TV ID is required",
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
	data, err := services.FetchContentUserData(supabaseClient, user, id, "tv")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch user TV details: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  data,
		"error": false,
	})
}
