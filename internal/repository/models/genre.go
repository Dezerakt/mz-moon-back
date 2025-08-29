package models

import (
	"mz-moon-back/internal/domain/catalog"

	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model

	Name string `gorm:"size:255"`
}

func (obj *Genre) ToEntity() catalog.Genre {
	return catalog.Genre{
		Name: obj.Name,
	}
}

func ToModel(entity *catalog.Genre) *Genre {
	return &Genre{
		Name: entity.Name,
	}
}
