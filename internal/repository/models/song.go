package models

import (
	"mz-moon-back/internal/domain/catalog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model

	UUID     uuid.UUID `gorm:"type:uuid;not null"`
	Name     string    `gorm:"size:255"`
	ArtistID uint      `gorm:"index"`
	GenreID  uint      `gorm:"index"`

	Genre  Genre  `gorm:"foreignKey:GenreID"`
	Artist Artist `gorm:"foreignKey:ArtistID"`
}

func (obj *Song) ToEntity() catalog.Song {
	return catalog.Song{
		UUID:     obj.UUID,
		SongName: obj.Name,
		Artist:   obj.Artist.Name,
		Genre:    obj.Genre.Name,
	}
}
