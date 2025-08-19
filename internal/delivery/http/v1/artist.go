package v1

import (
	"github.com/gofiber/fiber"
	"music-streaming/internal/domain/artist"
	"music-streaming/internal/usecase"
	"music-streaming/utils"
	errVo "music-streaming/utils/error"
)

type ArtistRouter struct {
	usecase usecase.IArtist
}

func NewArtistRouter(router fiber.Router, usecase usecase.IArtist) {
	r := ArtistRouter{
		usecase: usecase,
	}

	artistGroup := router.Group("/artist")
	{
		artistGroup.Post("/", r.create)
	}
}

func (obj *ArtistRouter) create(c *fiber.Ctx) {
	var (
		ctx     = c.Context()
		request = artist.Request{}
	)

	if err := c.BodyParser(&request); err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	if err := validate.Struct(&request); err != nil {
		utils.ReturnError(c, errVo.ValidateParamsError, err)
		return
	}

	create, err := obj.usecase.Create(ctx, request.ToEntity())
	if err != nil {
		utils.ReturnError(c, errVo.InternalError, err)
		return
	}

	utils.ReturnOk(c, create)
}
