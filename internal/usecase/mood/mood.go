package mood

import (
	moodDmn "accord-generator/internal/domain/mood"
	"accord-generator/internal/repository"
	"accord-generator/internal/usecase"
	"context"
	"github.com/pkg/errors"
)

type Usecase struct {
	MoodRepo repository.Mood
}

func NewMoodUsecase(moodRepo repository.Mood) usecase.Mood {
	return &Usecase{
		MoodRepo: moodRepo,
	}
}

func (obj *Usecase) GetMoods(ctx context.Context) ([]moodDmn.Mood, error) {
	result, err := obj.MoodRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Mood Usecase.GetMoods: ")
	}

	return result, nil
}
