package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cover struct {
	gorm.Model

	UUID uuid.UUID `gorm:"type:uuid"`
	Path string    `gorm:"size:255"`
}
