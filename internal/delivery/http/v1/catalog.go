package v1

import (
	"mz-moon-back/internal/delivery/http/v1/request"
	"mz-moon-back/internal/delivery/http/v1/response"
	"mz-moon-back/internal/usecase"
	errVo "mz-moon-back/utils/error"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

type CatalogRouter struct {
	u         usecase.ICatalog
	validator *validator.Validate
}

func NewCatalogRouter(router fiber.Router, usecase usecase.ICatalog) {
	r := CatalogRouter{
		u:         usecase,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	catalogRouter := router.Group("/catalog")
	{
		songRouter := catalogRouter.Group("/song")
		{
			songRouter.Get("/", r.getAllSongs)
		}

		genreRouter := catalogRouter.Group("/genre")
		{
			genreRouter.Get("/", r.getAllGenres)
			genreRouter.Put("/", r.newGenre)
		}
	}
}

func (obj *CatalogRouter) newGenre(c *fiber.Ctx) {
	var (
		ctx  = c.Context()
		body request.NewGenre
	)

	if err := c.BodyParser(&body); err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	if err := obj.validator.StructCtx(ctx, &body); err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	err := obj.u.NewGenre(ctx, body.ToEntity())
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	response.Ok(c, nil)
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
			UUID: genreEl.UUID,
			ID:   genreEl.ID,
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
