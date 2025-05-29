package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	jwtSecret := []byte(os.Getenv("SUPABASE_JWT_SECRET"))

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Next()
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		tokenStr = strings.TrimSpace(tokenStr)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return jwtSecret, nil
		})
		if err != nil {
			log.Printf("JWT parse error: %v", err)
			return c.Next()
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if sub, ok := claims["sub"].(string); ok {
				c.Locals("user_id", sub)
			}
			c.Locals("token", tokenStr)
		} else {
			log.Println("Invalid token or claims")
		}

		return c.Next()
	}
}
