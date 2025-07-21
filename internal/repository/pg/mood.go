package pgRepo

import (
	moodDmn "accord-generator/internal/domain/mood"
	"accord-generator/internal/repository"
	pgPkg "accord-generator/pkg/pg"
	"context"
)

type moodRepo struct {
	DB *pgPkg.PgWrap
}

func NewMoodRepo(pg *pgPkg.PgWrap) repository.Mood {
	return &moodRepo{
		DB: pg,
	}
}

func (obj *moodRepo) GetAll(ctx context.Context) ([]moodDmn.Mood, error) {
	return nil, nil
}
