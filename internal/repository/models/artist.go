package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model

	UUID uuid.UUID `gorm:"type:uuid;not null"`
	Name string    `gorm:"size:255"`
}
