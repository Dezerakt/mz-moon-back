package httpDlv

import (
	"mz-moon-back/config"
	v1 "mz-moon-back/internal/delivery/http/v1"
	"mz-moon-back/internal/usecase"

	"github.com/gofiber/fiber"
)

func NewRouter(app *fiber.App, cfg *config.Config, catalog usecase.ICatalog, media usecase.IMedia) {
	apiGroup := app.Group("/api")

	v1Group := apiGroup.Group("/v1")
	{
		v1.NewCatalogRouter(v1Group, catalog)
		v1.NewMediaRouter(v1Group, media)
	}

	app.Use(func(c *fiber.Ctx) {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "Route not found 😢",
		})
	})
}
