package songRepo

import (
	"context"
	"music-streaming/internal/domain/song"
	"music-streaming/internal/repository"
	pgPkg "music-streaming/pkg/pg"
)

type SongRepo struct {
	*pgPkg.PgWrap
}

func NewSongRepo(pgWrap *pgPkg.PgWrap) repository.ISong {
	return &SongRepo{
		PgWrap: pgWrap,
	}
}

func (obj *SongRepo) CreateNewMeta(ctx context.Context, songEntity *song.Entity, songPath string) (uint, error) {
	songModel := ToModel(songEntity)

	songModel.FilePath = songPath

	err := obj.Db.Create(songModel).Error
	if err != nil {
		return 0, err
	}

	return songModel.ID, err
}
