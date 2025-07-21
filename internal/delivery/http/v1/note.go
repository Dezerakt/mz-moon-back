package v1

import (
	"accord-generator/internal/usecase"
	"github.com/gofiber/fiber"
	"net/http"
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
		fiberCtx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	fiberCtx.Status(http.StatusOK).JSON(notes)
}
