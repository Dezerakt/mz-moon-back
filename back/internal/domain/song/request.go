package song

type Request struct {
	SongName string `json:"songName,omitempty" validate:"required"`
	ArtistID int    `json:"artistID,omitempty" validate:"required"`
	GenreID  int    `json:"genreID,omitempty" validate:"required"`
}

func (obj *Request) ToEntity() *Entity {
	return &Entity{
		SongName: obj.SongName,
		ArtistID: obj.ArtistID,
		GenreID:  obj.GenreID,
	}
}
