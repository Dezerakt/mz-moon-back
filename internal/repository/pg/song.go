package pgRepo

import (
	"spotifykiller/internal/repository"
	pgPkg "spotifykiller/pkg/pg"
)

type SongRepo struct {
	db *pgPkg.PgWrap
}

func NewSongRepo(pgWrap *pgPkg.PgWrap) repository.Song {
	return &SongRepo{
		db: pgWrap,
	}
}

func (obj *SongRepo) CreateNewMeta() {
}
