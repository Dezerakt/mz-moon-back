package repository

import (
	"context"
	"music-streaming/internal/domain/artist"
	"music-streaming/internal/domain/song"
)

type (
	ISong interface {
		CreateNewMeta(ctx context.Context, songEntity *song.Entity, songPath string) (uint, error)
	}

	IArtist interface {
		Create(ctx context.Context, artistEntity *artist.Entity) (uint, error)
	}
)
