package note

import (
	noteDmn "accord-generator/internal/domain/note"
	"accord-generator/internal/repository"
	"context"
)

type Usecase struct {
	noteRepo repository.Note
}

func NewUsecase(noteRepo repository.Note) *Usecase {
	return &Usecase{
		noteRepo: noteRepo,
	}
}

func (obj *Usecase) GetNotes(ctx context.Context) ([]noteDmn.Note, error) {
	result, err := obj.noteRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
