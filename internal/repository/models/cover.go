package models

import (
	"gorm.io/gorm"
)

type Cover struct {
	gorm.Model

	TypeID uint   `gorm:"index"`
	Path   string `gorm:"size:255"`
}
