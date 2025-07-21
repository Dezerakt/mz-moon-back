package httpDlv

import (
	"accord-generator/config"
	v1 "accord-generator/internal/delivery/http/v1"
	"accord-generator/internal/usecase"
	"github.com/gofiber/fiber"
)

func NewRouter(app *fiber.App, cfg *config.Config, container usecase.Dependencies) {
	apiV1Group := app.Group("/v1")
	{
		v1.NewMoodRouter(apiV1Group, container.Mood)
		v1.NewNoteRouter(apiV1Group, container.Note)
	}
}
