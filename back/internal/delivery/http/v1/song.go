package v1

import (
	"bytes"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"io"
	songDmn "music-streaming/internal/domain/song"
	"music-streaming/internal/usecase"
	"music-streaming/utils"
	errVo "music-streaming/utils/error"
	"strconv"
)

type SongRouter struct {
	usecase usecase.ISong
}

func NewSongRouter(router fiber.Router, usecase usecase.ISong) {
	r := SongRouter{
		usecase: usecase,
	}

	songGroup := router.Group("/song")
	{
		songGroup.Get("/:id", r.stream)
		songGroup.Post("/upload", r.upload)
		songGroup.Get("/", r.songs)
	}
}

var validate = validator.New()

func (obj *SongRouter) upload(c *fiber.Ctx) {
	var (
		ctx         = c.Context()
		songRequest = songDmn.Request{}
	)

	if err := c.BodyParser(&songRequest); err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	if err := validate.Struct(songRequest); err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	formFile, err := c.FormFile("file")
	if err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	formFileData, err := formFile.Open()
	if err != nil {
		utils.ReturnError(c, errVo.BadRequestError, err)
		return
	}

	defer formFileData.Close()

	buffer := &bytes.Buffer{}
	_, err = io.Copy(buffer, formFileData)
	if err != nil {
		utils.ReturnError(c, errVo.BadRequestError, err)
		return
	}

	upload, err := obj.usecase.Upload(ctx, songRequest.ToEntity(), buffer)
	if err != nil {
		utils.ReturnError(c, errVo.InternalError, err)
		return
	}

	utils.ReturnOk(c, upload)
}

func (obj *SongRouter) stream(c *fiber.Ctx) {
	var (
		ctx    = c.Context()
		songId = c.Params("id")
	)

	if songId == "" {
		utils.ReturnError(c, errVo.InvalidParams, errors.New("field `id` not found"))
		return
	}

	songIdInt, err := strconv.ParseUint(songId, 10, 64)
	if err != nil {
		utils.ReturnError(c, errVo.InvalidParams, err)
		return
	}

	filepath, err := obj.usecase.GetSongData(ctx, uint(songIdInt))
	if err != nil {
		return
	}

	c.Set("Content-Type", "audio/mpeg")
	c.Set("Accept-Ranges", "bytes")

	c.SendFile(filepath, true)
}

func (obj *SongRouter) songs(c *fiber.Ctx) {
	var (
		ctx = c.Context()
	)

	songs, err := obj.usecase.GetAllSongs(ctx)
	if err != nil {
		utils.ReturnError(c, errVo.BadRequestError, err)
		return
	}

	utils.ReturnOk(c, songs)
}
