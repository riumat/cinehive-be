package usercontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func AddUserTv(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	client := config.NewSupabaseClient(token)

	contentId := c.Params("id")

	code, message := services.AddUserContent(client, userID, contentId, "tv")

	if message != nil || code != 201 {
		return c.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   message.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Content added successfully",
	})
}

func AddUserTvWatchlist(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	client := config.NewSupabaseClient(token)

	contentId := c.Params("id")

	code, message := services.AddUserContentWatchlist(client, userID, contentId, "tv")

	if message != nil || code != 201 {
		return c.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   message.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Content added successfully",
	})
}

func EditUserTv(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	contentId := c.Params("id")

	client := config.NewSupabaseClient(token)

	type Input struct {
		Rating float64 `json:"rating" validate:"required"`
	}

	var data Input
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}

	code, message := services.EditRating(client, userID, contentId, "tv", data.Rating)

	if message != nil {
		return c.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   message.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Content edited successfully",
	})
}

func DeleteUserTv(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	client := config.NewSupabaseClient(token)

	contentId := c.Params("id")

	code, message := services.DeleteUserContent(client, userID, contentId, "tv")

	if message != nil {
		return c.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   message.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Content removed successfully",
	})
}

func DeleteUserTvWatchlist(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	client := config.NewSupabaseClient(token)

	contentId := c.Params("id")

	code, message := services.DeleteUserContentWatchlist(client, userID, contentId, "tv")

	if message != nil {
		return c.Status(code).JSON(fiber.Map{
			"error": true,
			"msg":   message.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Content removed successfully",
	})
}
