package artistRepo

import (
	"gorm.io/gorm"
	"music-streaming/internal/domain/artist"
)

type Artist struct {
	gorm.Model
	AristName string
}

func ToModel(artistEntity *artist.Entity) *Artist {
	return &Artist{
		AristName: artistEntity.ArtistName,
	}
}
