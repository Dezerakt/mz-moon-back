package progression

import (
	progressionDmn "accord-generator/internal/domain/progression"
	"accord-generator/internal/repository"
	"context"
	"github.com/pkg/errors"
)

type Usecase struct {
	progressionRepo repository.Progression
}

func NewUsecase(progressionRepo repository.Progression) *Usecase {
	return &Usecase{
		progressionRepo: progressionRepo,
	}
}

func (obj *Usecase) GetProgression(ctx context.Context, noteId, moodId string) ([]progressionDmn.Progression, error) {
	result, err := obj.progressionRepo.GetByParams(ctx, map[string]interface{}{
		"rootNote": noteId,
		"mood":     moodId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Progression Usecase. GetByParams: ")
	}

	return result, nil
}
