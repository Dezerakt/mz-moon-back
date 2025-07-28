package usecase

import (
	"context"
	"music-streaming/internal/domain/artist"
	"music-streaming/internal/repository"
)

type Artist struct {
	repository repository.IArtist
}

func NewArtistUsecase(repository repository.IArtist) IArtist {
	return &Artist{
		repository: repository,
	}
}

func (obj *Artist) Create(ctx context.Context, artistEntity *artist.Entity) (uint, error) {
	insertId, err := obj.repository.Create(ctx, artistEntity)
	if err != nil {
		return 0, err
	}

	return insertId, nil
}
