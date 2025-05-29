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
	tmdbClient := config.NewTMDBClient()
	data, err := services.FetchMovieDetails(tmdbClient, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token := c.Locals("token")

	if token != nil {
		user := c.Locals("user_id").(string)
		supabaseClient := config.NewSupabaseClient(token.(string))
		userData, err := services.FetchContentUserData(supabaseClient, user, data.Id, "movie")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "Failed to fetch user data" + err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":      data,
			"user_data": userData,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  data,
	})
}
