package usercontrollers

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/riumat/cinehive-be/app/services"
	"github.com/riumat/cinehive-be/config"
	"golang.org/x/sync/errgroup"
)

func GetUserWatchlist(c *fiber.Ctx) error {
	token := c.Locals("token")
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "Unauthorized",
		})
	}

	page := c.Query("page", "1")
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}

	user := c.Locals("user_id").(string)
	supabaseClient := config.NewSupabaseClient(token.(string))
	tmdbClient := config.NewTMDBClient()

	watchlistData, err := services.FetchWatchlist(supabaseClient, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch watchlist: " + err.Error(),
		})
	}

	limit := 10
	offset := (pageNum - 1) * limit
	totalItems := len(watchlistData)

	if offset >= totalItems {
		return c.JSON(fiber.Map{
			"data":      []interface{}{},
			"error":     false,
			"page":      pageNum,
			"totalPage": 0,
		})
	}

	end := offset + limit
	if end > totalItems {
		end = totalItems
	}

	paginatedData := watchlistData[offset:end]

	for i, item := range paginatedData {
		log.Printf("Item %d: %+v (type: %T)", i, item, item)
	}

	var detailedResults []interface{}
	var mu sync.Mutex
	g, _ := errgroup.WithContext(context.Background())

	for _, item := range paginatedData {
		itemMap := item

		contentID, ok := itemMap["content_id"].(float64)
		if !ok {
			continue
		}

		contentType, ok := itemMap["content_type"].(string)
		if !ok {
			continue
		}

		g.Go(func(id float64, mediaType string) func() error {
			return func() error {
				var details interface{}

				if mediaType == "movie" {
					movieDetails, err := services.FetchMovieDetails(tmdbClient, strconv.FormatFloat(contentID, 'f', 0, 64))
					if err != nil {
						return err
					}
					movieDetails.MediaType = "movie"

					details = movieDetails
				} else if mediaType == "tv" {
					tvDetails, err := services.FetchTvDetails(tmdbClient, strconv.FormatFloat(contentID, 'f', 0, 64))
					if err != nil {
						return err
					}
					tvDetails.MediaType = "tv"
					details = tvDetails

				} else {
					return nil
				}

				mu.Lock()
				detailedResults = append(detailedResults, details)
				mu.Unlock()
				return nil
			}
		}(contentID, contentType))
	}

	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to fetch content details: " + err.Error(),
		})
	}

	totalPages := (totalItems + limit - 1) / limit

	return c.JSON(fiber.Map{
		"data":      detailedResults,
		"error":     false,
		"page":      pageNum,
		"totalPage": totalPages,
	})
}
