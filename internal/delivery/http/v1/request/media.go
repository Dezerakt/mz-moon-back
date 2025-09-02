package request

import (
	"io"
	"mz-moon-back/internal/domain/media"

	"github.com/google/uuid"
)

type UpdateCover struct {
	ContentType media.ContentType `json:"content_type" validate:"required"`
	ForeignUUID string            `json:"foreign_id" validate:"required"`
	CoverFile   io.Reader         `json:"cover_file" validate:"required"`
}

func (obj *UpdateCover) ToEntity() *media.Cover {
	return &media.Cover{
		ContentType: obj.ContentType,
		ForeignUUID: uuid.MustParse(obj.ForeignUUID),
	}
}
