package repository

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
)

type (
	ISong interface {
		GetAll(ctx context.Context) ([]catalog.Song, error)
	}

	IGenre interface {
	}

	IArtist interface {
	}
)
