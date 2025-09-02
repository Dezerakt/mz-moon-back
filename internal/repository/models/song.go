package models

import (
	"mz-moon-back/internal/domain/catalog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model

	Name     string    `gorm:"size:255"`
	ArtistID uint      `gorm:"index"`
	GenreID  uint      `gorm:"index"`
	UUID     uuid.UUID `gorm:"index"`

	Genre  Genre  `gorm:"foreignKey:GenreID"`
	Artist Artist `gorm:"foreignKey:ArtistID"`
}

func (obj *Song) ToEntity() catalog.Song {
	return catalog.Song{
		SongName: obj.Name,
		Artist:   obj.Artist.Name,
		Genre:    obj.Genre.Name,
		UUID:     obj.UUID,
	}
}
