package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
)

func AddUserContent(c *fiber.Ctx) error {
	tokenStr, ok := c.Locals("token").(string)
	if !ok || tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	// Decodifica il token JWT e ottieni il sub
	jwtSecret := []byte(os.Getenv("SUPABASE_JWT_SECRET"))
	token, message := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if message != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: invalid token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: invalid claims",
		})
	}
	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing user id",
		})
	}

	client := config.NewSupabaseClient(tokenStr)

	type Input struct {
		ContentId   float64 `json:"content_id" validate:"required"`
		ContentType string  `json:"content_type" validate:"required,oneof=movie tv"`
	}

	var data Input
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}

	code, message := services.AddUserContent(client, userID, data.ContentId, data.ContentType)

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
