package httpDlv

import (
	"github.com/gofiber/fiber"
	"spotifykiller/config"
	v1 "spotifykiller/internal/delivery/http/v1"
	"spotifykiller/internal/usecase"
)

func NewRouter(app *fiber.App, cfg *config.Config, song usecase.Song) {
	apiV1Group := app.Group("/v1")
	{
		v1.NewSongRouter(apiV1Group, song)
	}
}
