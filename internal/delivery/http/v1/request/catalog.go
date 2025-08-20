package request

type SongRequest struct {
	SongName string `json:"songName,omitempty" validate:"required"`
	ArtistID uint   `json:"artistID,omitempty" validate:"required"`
	GenreID  uint   `json:"genreID,omitempty" validate:"required"`
}
