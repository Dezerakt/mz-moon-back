package song

type Request struct {
	SongName string `json:"songName,omitempty" validate:"required"`
	ArtistID uint   `json:"artistID,omitempty" validate:"required"`
	GenreID  uint   `json:"genreID,omitempty" validate:"required"`
}

func (obj *Request) ToEntity() *Entity {
	return &Entity{
		SongName: obj.SongName,
		ArtistID: obj.ArtistID,
		GenreID:  obj.GenreID,
	}
}
