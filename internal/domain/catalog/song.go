package catalog

type Song struct {
	SongName string
	ArtistID uint
	GenreID  uint
	FilePath string
}

type SongRequest struct {
	SongName string `json:"songName,omitempty" validate:"required"`
	ArtistID uint   `json:"artistID,omitempty" validate:"required"`
	GenreID  uint   `json:"genreID,omitempty" validate:"required"`
}

func (obj *SongRequest) ToEntity() *Song {
	return &Song{
		SongName: obj.SongName,
		ArtistID: obj.ArtistID,
		GenreID:  obj.GenreID,
	}
}
