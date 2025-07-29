package songRepo

import (
	"gorm.io/gorm"
	"music-streaming/internal/domain/song"
	"music-streaming/internal/repository/artist"
	"music-streaming/internal/repository/genre"
)

type Song struct {
	gorm.Model
	ArtistID uint
	GenreID  uint
	Name     string
	FilePath string

	Artist artistRepo.Artist `gorm:"foreignKey:ArtistID"`
	Genre  genreRepo.Genre   `gorm:"foreignKey:GenreID"`
}

func ToModel(entity *song.Entity) *Song {
	return &Song{
		ArtistID: uint(entity.ArtistID),
		GenreID:  uint(entity.GenreID),
		Name:     entity.SongName,
	}
}

func ToEntity(model *Song) *song.Entity {
	return &song.Entity{
		SongName: model.Name,
		ArtistID: model.ArtistID,
		GenreID:  model.GenreID,
		FilePath: model.FilePath,
	}
}
