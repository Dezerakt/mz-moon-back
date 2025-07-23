package v1

import (
	progressionDmn "accord-generator/internal/domain/progression"
	"accord-generator/internal/usecase"
	"accord-generator/utils"
	"accord-generator/utils/error"
	"github.com/gofiber/fiber"
)

type ProgressionRouter struct {
	progressionUsecase usecase.Progression
}

func NewProgressionRouter(router fiber.Router, progression usecase.Progression) {
	r := ProgressionRouter{
		progressionUsecase: progression,
	}

	progressionGroup := router.Group("/progression")
	{
		progressionGroup.Get("/", r.getProgression)
	}
}

func (obj *ProgressionRouter) getProgression(fiberCtx *fiber.Ctx) {
	var (
		ctx     = fiberCtx.Context()
		request progressionDmn.Request
	)

	if err := fiberCtx.QueryParser(&request); err != nil {
		utils.ReturnError(fiberCtx, errVo.BadRequestError, err)
		return
	}

	progression, err := obj.progressionUsecase.GetProgression(ctx, request.RootNote, request.Mood)
	if err != nil {
		utils.ReturnError(fiberCtx, errVo.InternalError, err)
		return
	}

	utils.ReturnOk(fiberCtx, progression)
}
