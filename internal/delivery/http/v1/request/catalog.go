package request

import (
	"mz-moon-back/internal/domain/catalog"
)

type NewGenre struct {
	Name string `json:"name" form:"name" validate:"required"`
}

func (req *NewGenre) ToEntity() *catalog.Genre {
	return &catalog.Genre{
		Name: req.Name,
	}
}
