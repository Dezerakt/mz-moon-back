package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cover struct {
	gorm.Model

	ContentType string    `gorm:"index"`
	ForeignUUID uuid.UUID `gorm:"type:uuid;default:null"`
}
