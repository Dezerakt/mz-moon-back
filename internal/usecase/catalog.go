package usecase

import (
	"bytes"
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/domain/media"
	"mz-moon-back/internal/repository"

	"github.com/google/uuid"
)

type Catalog struct {
	songRepo    repository.ISong
	artistRepo  repository.IArtist
	genreRepo   repository.IGenre
	coverRepo   repository.ICover
	fileManager repository.IFileManager
}

func NewCatalog(songRepo repository.ISong,
	artistRepo repository.IArtist,
	genreRepo repository.IGenre,
	coverRepo repository.ICover,
	fileManager repository.IFileManager) ICatalog {
	return &Catalog{
		songRepo:    songRepo,
		artistRepo:  artistRepo,
		genreRepo:   genreRepo,
		coverRepo:   coverRepo,
		fileManager: fileManager,
	}
}

func (obj *Catalog) GetSongs(ctx context.Context) ([]catalog.Song, error) {
	songs, err := obj.songRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (obj *Catalog) GetGenres(ctx context.Context) ([]catalog.Genre, error) {
	genres, err := obj.genreRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (obj *Catalog) NewGenre(ctx context.Context, genre *catalog.Genre, buffer *bytes.Buffer) error {
	newGenreUUID := uuid.New()

	genre.UUID = newGenreUUID

	err := obj.genreRepo.NewGenre(ctx, genre)
	if err != nil {
		return err
	}

	err = obj.coverRepo.UpsertCover(ctx, &media.Cover{
		ContentType: media.GenreCover,
		ForeignUUID: newGenreUUID,
	})
	if err != nil {
		return err
	}

	err = obj.fileManager.StoreCover(ctx, newGenreUUID, media.GenreCover, buffer)
	if err != nil {
		return err
	}

	return nil
}
