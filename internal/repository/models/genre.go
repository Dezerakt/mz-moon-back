package models

import (
	"mz-moon-back/internal/domain/catalog"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model

	UUID uuid.UUID `gorm:"type:uuid;not null"`
	Name string    `gorm:"size:255"`
}

func (obj *Genre) ToEntity() catalog.Genre {
	return catalog.Genre{
		UUID: obj.UUID,
		Name: obj.Name,
	}
}
