package v1

import (
	"accord-generator/internal/usecase"
	"accord-generator/utils"
	errVo "accord-generator/utils/error"
	"github.com/gofiber/fiber"
)

type MoodRouter struct {
	moodUsecase usecase.Mood
}

func NewMoodRouter(router fiber.Router, mood usecase.Mood) {
	r := MoodRouter{
		moodUsecase: mood,
	}

	moodGroup := router.Group("/mood")
	{
		moodGroup.Get("/", r.getMoods)
	}
}

func (obj *MoodRouter) getMoods(fiberCtx *fiber.Ctx) {
	var (
		ctx = fiberCtx.Context()
	)

	moods, err := obj.moodUsecase.GetMoods(ctx)
	if err != nil {
		utils.ReturnError(fiberCtx, errVo.BadRequestError, err)
		return
	}

	utils.ReturnOk(fiberCtx, moods)
}
