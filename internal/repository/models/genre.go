package models

import (
	"mz-moon-back/internal/domain/catalog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model

	Name     string    `gorm:"size:255"`
	WebTitle string    `gorm:"size:255"`
	UUID     uuid.UUID `gorm:"type:uuid;default:null"`
}

func (obj *Genre) ToEntity() catalog.Genre {
	return catalog.Genre{
		Name:     obj.Name,
		UUID:     obj.UUID,
		ID:       obj.ID,
		WebTitle: obj.WebTitle,
	}
}

func ToModel(entity *catalog.Genre) *Genre {
	return &Genre{
		Name:     entity.Name,
		UUID:     entity.UUID,
		WebTitle: entity.WebTitle,
	}
}
