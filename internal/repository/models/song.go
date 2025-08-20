package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model

	Name     string `gorm:"size:255"`
	ArtistID uint   `gorm:"index"`
	GenreID  uint   `gorm:"index"`

	Genre  Genre  `gorm:"foreignKey:GenreID"`
	Artist Artist `gorm:"foreignKey:ArtistID"`
}
