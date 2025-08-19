package artistRepo

import (
	"context"
	"music-streaming/internal/domain/artist"
	"music-streaming/internal/repository"
	pgPkg "music-streaming/pkg/pg"
)

type ArtistRepo struct {
	*pgPkg.PgWrap
}

func NewArtistRepo(pgWrap *pgPkg.PgWrap) repository.IArtist {
	return &ArtistRepo{
		PgWrap: pgWrap,
	}
}

func (obj *ArtistRepo) Create(ctx context.Context, songEntity *artist.Entity) (uint, error) {
	artistModel := ToModel(songEntity)

	err := obj.Db.Create(artistModel).Error
	if err != nil {
		return 0, err
	}

	return artistModel.ID, err
}
