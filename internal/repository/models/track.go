package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model

	UUID uuid.UUID `gorm:"type:uuid"`
}
