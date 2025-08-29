package media

type ContentTypeID uint

const (
	SongCover ContentTypeID = iota
	GenreCover
	ProfilePicture
)

type Cover struct {
	ContentTypeID ContentTypeID
	ForeignID     uint
	Path          string
}
