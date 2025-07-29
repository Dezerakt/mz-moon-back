package usecase

import (
	"bytes"
	"context"
	"music-streaming/internal/domain/artist"
	"music-streaming/internal/domain/song"
)

type (
	ISong interface {
		Upload(ctx context.Context, songEntity *song.Entity, buffer *bytes.Buffer) (uint, error)
		GetSongData(ctx context.Context, songId uint) (string, error)
	}

	IArtist interface {
		Create(ctx context.Context, artistEntity *artist.Entity) (uint, error)
	}
)
