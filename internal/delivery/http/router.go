package httpDlv

import (
	"github.com/gofiber/fiber"
	"mz-moon-back/config"
	v1 "mz-moon-back/internal/delivery/http/v1"
	"mz-moon-back/internal/usecase"
)

func NewRouter(app *fiber.App, cfg *config.Config, catalog usecase.ICatalog) {
	apiGroup := app.Group("/api")

	v1Group := apiGroup.Group("/v1")
	{
		v1.NewCatalogRouter(v1Group, catalog)
	}

	app.Use(func(c *fiber.Ctx) {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "Route not found 😢",
		})
	})
}
