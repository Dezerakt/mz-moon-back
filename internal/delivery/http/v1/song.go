package v1

import (
	"bytes"
	"github.com/gofiber/fiber"
	"io"
	songDmn "spotifykiller/internal/domain/song"
	"spotifykiller/internal/usecase"
	"spotifykiller/utils"
	errVo "spotifykiller/utils/error"
)

type SongRouter struct {
	usecase usecase.Song
}

func NewSongRouter(router fiber.Router, usecase usecase.Song) {
	r := SongRouter{
		usecase: usecase,
	}

	songGroup := router.Group("/usecase")
	{
		songGroup.Get("/:id", r.stream)
		songGroup.Post("/upload", r.upload)
	}
}

func (obj *SongRouter) upload(c *fiber.Ctx) {
	var (
		ctx         = c.Context()
		songRequest = &songDmn.Request{}
	)

	if err := c.BodyParser(songRequest); err != nil {
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

	obj.usecase.Upload(ctx, songRequest, buffer)
}

func (obj *SongRouter) stream(c *fiber.Ctx) {

}
