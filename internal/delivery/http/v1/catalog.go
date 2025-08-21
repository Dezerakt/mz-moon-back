package v1

import (
	"mz-moon-back/internal/delivery/http/v1/response"
	"mz-moon-back/internal/usecase"
	errVo "mz-moon-back/utils/error"

	"github.com/gofiber/fiber"
)

type CatalogRouter struct {
	u usecase.ICatalog
}

func NewCatalogRouter(router fiber.Router, usecase usecase.ICatalog) {
	r := CatalogRouter{
		u: usecase,
	}

	catalogRouter := router.Group("/catalog")
	{
		catalogRouter.Get("/song", r.getAllSongs)
	}
}

func (obj *CatalogRouter) getAllSongs(c *fiber.Ctx) {
	var (
		ctx = c.Context()
	)

	entitySongs, err := obj.u.GetSongs(ctx)
	if err != nil {
		response.Error(c, errVo.InvalidParams, err)
		return
	}

	var responseSongs []response.GetAllSongs
	for _, entitySong := range entitySongs {
		responseSongs = append(responseSongs, response.GetAllSongs{
			Song:  entitySong.SongName,
			Arist: entitySong.Artist,
			Genre: entitySong.Genre,
		})
	}

	response.Ok(c, responseSongs)
}
