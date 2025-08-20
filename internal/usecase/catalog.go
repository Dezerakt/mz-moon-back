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

func NewCatalog(songRepo repository.ISong, artistRepo repository.IArtist, genreRepo repository.IGenre) ICatalog {
	return &Catalog{
		songRepo:   songRepo,
		artistRepo: artistRepo,
		genreRepo:  genreRepo,
	}
}

func (obj *Catalog) GetSongs(ctx context.Context) ([]catalog.Song, error) {
	songs, err := obj.songRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
