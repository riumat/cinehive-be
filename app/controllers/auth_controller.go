package controllers

import (
	/* "context" */

	"time"

	/* "github.com/create-go-app/fiber-go-template/platform/cache" */
	"github.com/riumat/cinehive-be/app/models"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/database"
	"github.com/riumat/cinehive-be/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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
	if err := services.AddUserToProfile(data.User.ID, input.Username); err != nil {
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

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	// Create a new user auth struct.
	signUp := &models.SignUp{}
	database.ConnectDB()

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	/* 	db, err := database.OpenDBConnection()
	   	if err != nil {
	   		// Return status 500 and database connection error.
	   		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	   			"error": true,
	   			"msg":   err.Error(),
	   		})
	   	}
	*/
	// Checking role from sign up data.
	/* role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	} */

	// Create a new user struct.
	user := &models.User{}

	// Set initialized default data for user:
	user.UserID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = signUp.Email
	user.Username = signUp.Username
	user.Password = utils.GeneratePassword(signUp.Password)

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create a new user with validated data.
	if err := database.DB.Create(user).Error; err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Delete password hash field from JSON view.
	user.Password = ""

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// UserSignIn method to auth user and return access and refresh tokens.
// @Description Auth user and return access and refresh token.
// @Summary auth user and return access and refresh token
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Success 200 {string} status "ok"
// @Router /v1/user/sign/in [post]
func UserSignIn(c *fiber.Ctx) error {
	// Struttura per i dati di input del login
	signIn := &models.SignIn{}

	// Parsing del corpo della richiesta JSON
	if err := c.BodyParser(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid input data",
		})
	}

	// Validazione dei campi di input
	validate := utils.NewValidator()
	if err := validate.Struct(signIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Trova l'utente nel database tramite l'email
	var user models.User
	if err := database.DB.Where("email = ?", signIn.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid email or password",
		})
	}

	// Confronta la password fornita con quella salvata nel database
	if !utils.ComparePasswords(user.Password, signIn.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid email or password",
		})
	}

	// Genera un nuovo token di accesso (JWT) e un token di refresh
	tokens, err := utils.GenerateNewTokens(user.UserID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to generate tokens",
		})
	}

	// Restituisci i token generati
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Login successful",
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}

// UserSignOut method to de-authorize user and delete refresh token from Redis.
// @Description De-authorize user and delete refresh token from Redis.
// @Summary de-authorize user and delete refresh token from Redis
// @Tags User
// @Accept json
// @Produce json
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user/sign/out [post]
/* func UserSignOut(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Define user ID.
	userID := claims.UserID.String()

	// Create a new Redis connection.
	connRedis, err := cache.RedisConnection()
	if err != nil {
		// Return status 500 and Redis connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Save refresh token to Redis.
	errDelFromRedis := connRedis.Del(context.Background(), userID).Err()
	if errDelFromRedis != nil {
		// Return status 500 and Redis deletion error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   errDelFromRedis.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
} */
