package usecase

import (
	"bytes"
	"context"
	"mz-moon-back/internal/domain/catalog"
	"os"

	"github.com/google/uuid"
)

type (
	ICatalog interface {
		GetSongs(ctx context.Context) ([]catalog.Song, error)
		GetGenres(ctx context.Context) ([]catalog.Genre, error)
	}

	IMedia interface {
		GetTrack(ctx context.Context, songUUID uuid.UUID) (*os.File, error)
		GetTrackCover(ctx context.Context, songUUID uuid.UUID) (*bytes.Buffer, error)
	}
)
