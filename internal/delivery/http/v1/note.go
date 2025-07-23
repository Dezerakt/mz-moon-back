package v1

import (
	"accord-generator/internal/usecase"
	"accord-generator/utils"
	errVo "accord-generator/utils/error"
	"github.com/gofiber/fiber"
)

type NoteRouter struct {
	noteUsecase usecase.Note
}

func NewNoteRouter(router fiber.Router, note usecase.Note) {
	r := NoteRouter{
		noteUsecase: note,
	}

	noteGroup := router.Group("/note")
	{
		noteGroup.Get("/", r.getNotes)
	}
}

func (obj *NoteRouter) getNotes(fiberCtx *fiber.Ctx) {
	var (
		ctx = fiberCtx.Context()
	)

	notes, err := obj.noteUsecase.GetNotes(ctx)
	if err != nil {
		utils.ReturnError(fiberCtx, errVo.BadRequestError, err)
		return
	}

	utils.ReturnOk(fiberCtx, notes)
}
