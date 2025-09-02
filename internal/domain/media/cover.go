package media

import "github.com/google/uuid"

type ContentType string

const (
	SongCover      ContentType = "song"
	GenreCover     ContentType = "genre"
	ProfilePicture ContentType = "profile"
)

var ContentTypeList = []ContentType{
	SongCover,
	GenreCover,
	ProfilePicture,
}

type Cover struct {
	ContentType ContentType
	ForeignUUID uuid.UUID
}
