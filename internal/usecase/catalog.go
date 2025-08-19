package usecase

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/repository"
)

type Catalog struct {
	songRepo   repository.ISong
	artistRepo repository.IArtist
	genreRepo  repository.IGenre
}

func (obj *Catalog) GetSongs(ctx context.Context) ([]catalog.Song, error) {

	return nil, nil
}
