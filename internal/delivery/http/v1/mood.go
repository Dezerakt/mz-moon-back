package v1

import (
	"accord-generator/internal/usecase"
	"github.com/gofiber/fiber"
	"net/http"
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
		fiberCtx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	fiberCtx.Status(http.StatusOK).JSON(moods)
}
