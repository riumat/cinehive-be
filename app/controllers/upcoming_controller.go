package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
	"golang.org/x/sync/errgroup"
)

func GetUpcoming(c *fiber.Ctx) error {
	client := config.NewTMDBClient()

	var movieData, tvData any
	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		data, err := services.FetchUpcoming(client, "movie")
		if err != nil {
			return err
		}
		movieData = data
		return nil
	})

	g.Go(func() error {
		data, err := services.FetchUpcoming(client, "tv")
		if err != nil {
			return err
		}
		tvData = data
		return nil
	})

	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"movies": movieData,
			"tv":     tvData,
		},
	})
}
