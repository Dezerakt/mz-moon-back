package repository

import (
	moodDmn "accord-generator/internal/domain/mood"
	noteDmn "accord-generator/internal/domain/note"
	progressionDmn "accord-generator/internal/domain/progression"
	"context"
)

type (
	Mood interface {
		GetAll(ctx context.Context) ([]moodDmn.Mood, error)
	}

	Note interface {
		GetAll(ctx context.Context) ([]noteDmn.Note, error)
	}

	Progression interface {
		GetAll(ctx context.Context) ([]progressionDmn.Progression, error)
		GetByParams(ctx context.Context, params map[string]interface{}) ([]progressionDmn.Progression, error)
	}
)
