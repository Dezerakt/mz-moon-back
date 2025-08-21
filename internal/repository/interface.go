package repository

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/domain/media"

	"github.com/google/uuid"
)

type (
	ISong interface {
		GetAll(ctx context.Context) ([]catalog.Song, error)
	}

	IGenre interface {
	}

	IArtist interface {
	}

	ITrack interface {
		GetTrack(ctx context.Context, uuid uuid.UUID) (*media.Track, error)
	}

	ICover interface {
		GetCover(ctx context.Context, uuid uuid.UUID) (*media.Cover, error)
	}
)
