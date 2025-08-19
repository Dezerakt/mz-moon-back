package usecase

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
)

type (
	ICatalog interface {
		GetSongs(ctx context.Context) ([]catalog.Song, error)
	}
)
