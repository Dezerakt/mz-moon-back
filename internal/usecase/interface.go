package usecase

import (
	moodDmn "accord-generator/internal/domain/mood"
	noteDmn "accord-generator/internal/domain/note"
	"context"
)

type (
	Mood interface {
		GetMoods(ctx context.Context) ([]moodDmn.Mood, error)
	}

	Note interface {
		GetNotes(ctx context.Context) ([]noteDmn.Note, error)
	}
)
