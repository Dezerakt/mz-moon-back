package repository

import (
	moodDmn "accord-generator/internal/domain/mood"
	noteDmn "accord-generator/internal/domain/note"
	"context"
)

type (
	Mood interface {
		GetAll(ctx context.Context) ([]moodDmn.Mood, error)
	}

	Note interface {
		GetAll(ctx context.Context) ([]noteDmn.Note, error)
	}
)
