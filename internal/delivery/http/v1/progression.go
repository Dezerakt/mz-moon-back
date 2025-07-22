package v1

import (
	progressionDmn "accord-generator/internal/domain/progression"
	"accord-generator/internal/usecase"
	"github.com/gofiber/fiber"
	"net/http"
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
		progressionGroup.Post("/", r.getProgression)
	}
}

func (obj *ProgressionRouter) getProgression(fiberCtx *fiber.Ctx) {
	var (
		ctx     = fiberCtx.Context()
		request progressionDmn.Request
	)

	if err := fiberCtx.BodyParser(&request); err != nil {
		fiberCtx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	progression, err := obj.progressionUsecase.GetProgression(ctx, request.NoteId, request.MoodId)
	if err != nil {
		fiberCtx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	fiberCtx.Status(http.StatusOK).JSON(progression)
}
