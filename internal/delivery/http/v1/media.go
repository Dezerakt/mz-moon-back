package v1

import (
	"bytes"
	"errors"
	"io"
	"mz-moon-back/internal/delivery/http/v1/request"
	"mz-moon-back/internal/delivery/http/v1/response"
	"mz-moon-back/internal/domain/media"
	"mz-moon-back/internal/usecase"
	errVo "mz-moon-back/utils/error"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

type MediaRouter struct {
	u         usecase.IMedia
	validator *validator.Validate
}

func NewMediaRouter(router fiber.Router, usecase usecase.IMedia) {
	r := MediaRouter{
		u:         usecase,
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	mediaRouter := router.Group("/media")
	{
		songGroup := mediaRouter.Group("/song")
		{
			songGroup.Get("/:uuid", r.streamSong)
		}

		coverGroup := mediaRouter.Group("/cover")
		{
			coverGroup.Get("/:uuid", r.getCover)
			coverGroup.Put("/", r.updateCover)
		}
	}
}

func (obj *MediaRouter) updateCover(c *fiber.Ctx) {
	var (
		ctx  = c.Context()
		body = request.UpdateCover{}
	)

	if err := c.BodyParser(&body); err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	err := obj.validator.StructCtx(ctx, &body)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	if !slices.Contains(media.ContentTypeList, body.ContentType) {
		response.Error(c, errVo.BadRequestError, errors.New("content type is not valid"))
		return
	}

	buffer := bytes.Buffer{}
	_, err = io.Copy(&buffer, body.CoverFile)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	err = obj.u.UpdateCover(ctx, body.ToEntity(), buffer)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
		return
	}

	response.Ok(c, nil)
}

func (obj *MediaRouter) getCover(c *fiber.Ctx) {
	var (
		ctx       = c.Context()
		queryUUID = c.Params("uuid")
	)

	songUUID, err := uuid.Parse(queryUUID)
	if err != nil {
		response.Error(c, errVo.InvalidParams, err)
		return
	}

	coverBytes, err := obj.u.GetTrackCover(ctx, songUUID)
	if err != nil {
		response.Error(c, errVo.BadRequestError, err)
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
