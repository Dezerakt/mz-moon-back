package v1

import (
	"context"
	"mz-moon-back/internal/delivery/http/v1/response"
	"mz-moon-back/internal/usecase"
	errVo "mz-moon-back/utils/error"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

type MediaRouter struct {
	u usecase.IMedia
}

func NewMediaRouter(router fiber.Router, usecase usecase.IMedia) {
	r := MediaRouter{u: usecase}

	mediaRouter := router.Group("/media")
	{
		songGroup := mediaRouter.Group("/song")
		{
			songGroup.Get("/:uuid", r.streamSong)
		}

		coverGroup := mediaRouter.Group("/cover")
		{
			coverGroup.Get("/:uuid", r.getCover)
			coverGroup.Put("/", r.newCover)
		}
	}
}

func (obj *MediaRouter) newCover(c *fiber.Ctx) {

}

func (obj *MediaRouter) getCover(c *fiber.Ctx) {
	var (
		ctx       = context.Background()
		queryUUID = c.Params("uuid")
	)

	songUUID, err := uuid.Parse(queryUUID)
	if err != nil {
		response.Error(c, errVo.InvalidParams, err)
		return
	}

	coverBytes, err := obj.u.GetTrackCover(ctx, songUUID)
	if err != nil {
		return
	}

	c.Set("Content-Type", "image/png")
	c.Send(coverBytes)
}

func (obj *MediaRouter) streamSong(c *fiber.Ctx) {
	var (
		ctx       = c.Context()
		queryUUID = c.Params("uuid")
	)

	songUUID, err := uuid.Parse(queryUUID)
	if err != nil {
		response.Error(c, errVo.InvalidParams, err)
		return
	}

	trackBytes, err := obj.u.GetTrack(ctx, songUUID)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	c.Set("Content-Type", "audio/mpeg")
	c.SendStream(trackBytes)
}
