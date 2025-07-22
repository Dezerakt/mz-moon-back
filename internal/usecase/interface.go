package usecase

import (
	moodDmn "accord-generator/internal/domain/mood"
	noteDmn "accord-generator/internal/domain/note"
	progressionDmn "accord-generator/internal/domain/progression"
	"context"
)

type (
	Mood interface {
		GetMoods(ctx context.Context) ([]moodDmn.Mood, error)
	}

	Note interface {
		GetNotes(ctx context.Context) ([]noteDmn.Note, error)
	}

	Progression interface {
		GetProgression(ctx context.Context, noteId, moodId string) ([]progressionDmn.Progression, error)
	}
)
