package request

import (
	"mz-moon-back/internal/domain/media"
)

type NewCover struct {
	ContentTypeID uint `json:"content_type_id"`
	ForeignID     uint `json:"foreign_id"`
}

func (obj *NewCover) ToEntity() *media.Cover {
	return &media.Cover{
		ContentTypeID: media.ContentTypeID(obj.ContentTypeID),
		ForeignID:     obj.ForeignID,
	}
}
