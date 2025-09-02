package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model

	Name string    `gorm:"size:255"`
	UUID uuid.UUID `gorm:"type:uuid;default:null"`
}
