package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model

	UUID     uuid.UUID `gorm:"type:uuid"`
	Name     string    `gorm:"size:255"`
	ArtistID uint      `gorm:"index"`
	GenreID  uint      `gorm:"index"`

	Genre  Genre  `gorm:"foreignKey:GenreID"`
	Artist Artist `gorm:"foreignKey:ArtistID"`
}
