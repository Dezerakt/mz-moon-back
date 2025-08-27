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
		catalogRouter.Get("/genre", r.getAllGenres)
	}
}

func (obj *CatalogRouter) getAllGenres(c *fiber.Ctx) {
	var (
		ctx = c.Context()
	)

	genres, err := obj.u.GetGenres(ctx)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	var resp []response.GetAllGenres
	for _, genreEl := range genres {
		resp = append(resp, response.GetAllGenres{
			Name: genreEl.Name,
			UUID: genreEl.UUID.String(),
		})
	}

	response.Ok(c, resp)
}

func (obj *CatalogRouter) getAllSongs(c *fiber.Ctx) {
	var (
		ctx = c.Context()
	)

	songs, err := obj.u.GetSongs(ctx)
	if err != nil {
		response.Error(c, errVo.InvalidParams, err)
		return
	}

	var resp []response.GetAllSongs
	for _, songsEl := range songs {
		resp = append(resp, response.GetAllSongs{
			Song:  songsEl.SongName,
			Arist: songsEl.Artist,
			Genre: songsEl.Genre,
			UUID:  songsEl.UUID,
		})
	}

	response.Ok(c, resp)
}
