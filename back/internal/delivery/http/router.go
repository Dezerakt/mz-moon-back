package httpDlv

import (
	"github.com/gofiber/fiber"
	"music-streaming/config"
	v1 "music-streaming/internal/delivery/http/v1"
	"music-streaming/internal/usecase"
)

func NewRouter(app *fiber.App, cfg *config.Config, song usecase.ISong, artist usecase.IArtist) {
	apiV1Group := app.Group("/v1")
	{
		v1.NewSongRouter(apiV1Group, song)
		v1.NewArtistRouter(apiV1Group, artist)
	}

	app.Use(func(c *fiber.Ctx) {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "Route not found 😢",
		})
	})
}
