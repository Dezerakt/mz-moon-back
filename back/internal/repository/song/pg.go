package songRepo

import (
	"context"
	"gorm.io/gorm"
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

func (obj *SongRepo) GetMetaById(ctx context.Context, songId uint) (*song.Entity, error) {
	songModel := Song{}

	err := obj.Db.Where(&Song{
		Model: gorm.Model{
			ID: songId,
		},
	}).First(&songModel).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&songModel), nil
}

func (obj *SongRepo) GetAllSongs(ctx context.Context) ([]song.Entity, error) {
	var modelSongs = make([]Song, 0)

	err := obj.Db.Find(&modelSongs).Error
	if err != nil {
		return nil, err
	}

	var songsEntity = make([]song.Entity, 0)
	for _, modelSong := range modelSongs {
		songsEntity = append(songsEntity, *ToEntity(&modelSong))
	}

	return songsEntity, nil
}
