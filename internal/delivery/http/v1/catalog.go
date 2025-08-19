package v1

import (
	"github.com/gofiber/fiber"
	"mz-moon-back/internal/usecase"
	"mz-moon-back/utils"
	errVo "mz-moon-back/utils/error"
)

type CatalogRouter struct {
	u usecase.ICatalog
}

func NewCatalogRouter(router fiber.Router, usecase usecase.ICatalog) {
	r := CatalogRouter{
		u: usecase,
	}

	catalogGroup := router.Group("/catalog")
	{
		catalogGroup.Get("/song", r.getSongs)
	}
}

func (obj *CatalogRouter) getSongs(c *fiber.Ctx) {
	var (
		ctx = c.Context()
	)

	songs, err := obj.u.GetSongs(ctx)
	if err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	utils.ReturnOk(c, songs)
}
