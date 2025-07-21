package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/pkg/utils"
)

func GetUserProfile(c *fiber.Ctx) error {
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized: missing or invalid token",
		})
	}

	profile, err := services.FetchMe(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch user profile: " + err.Error(),
		})
	}

	if profile.Code == 404 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "User profile not found",
		})
	}

	if profile.Code != 0 && profile.Code != 200 {
		return c.Status(profile.Code).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch user profile",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  profile.User,
	})
}

func SupabaseUserSignUp(c *fiber.Ctx) error {
	type SignUpInput struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,min=3,max=30"`
		Password string `json:"password" validate:"required,min=6"`
	}

	var input SignUpInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	data, err := services.SignUpWithSupabase(input.Email, input.Password)
	if err != nil {
		if data.Code == 409 {
			return c.Status(data.Code).JSON(fiber.Map{
				"error": true,
				"msg":   "User already exists",
			})
		}
		if data.Code == 422 {
			return c.Status(data.Code).JSON(fiber.Map{
				"error": true,
				"msg":   "Invalid email or already registered",
			})
		}
		if data.Code == 400 {
			return c.Status(data.Code).JSON(fiber.Map{
				"error": true,
				"msg":   "Invalid input data",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := services.AddUserToProfile(data.User.ID, input.Username, nil); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to add user to profile: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "User registered successfully",
		"data":  data,
	})
}

func SupabaseUserSignIn(c *fiber.Ctx) error {
	type SignInInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	var input SignInInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}
	data, err := services.SignInWithSupabase(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if data.Code == 400 {
		return c.Status(data.Code).JSON(fiber.Map{
			"error": true,
			"msg":   "Wrong credentials",
		})
	}
	return c.Status(data.Code).JSON(fiber.Map{
		"error": false,
		"data":  data,
	})
}
