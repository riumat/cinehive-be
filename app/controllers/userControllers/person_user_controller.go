package usercontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func AddUserPerson(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	userID := c.Locals("user_id").(string)

	client := config.NewSupabaseClient(token)

	id := c.Params("id")

	var input struct {
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}

	personInfo := services.PersonInfo{
		Name:        input.Name,
		Id:          id,
		ProfilePath: input.ProfilePath,
	}

	code, message := services.AddPerson(client, userID, personInfo)

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

func DeleteUserPerson(c *fiber.Ctx) error {
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

	code, message := services.DeletePerson(client, userID, contentId)

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
