package response

type GetAllSongs struct {
	Song  string `json:"name,omitempty" validate:"required"`
	Arist string `json:"artist,omitempty" validate:"required"`
	Genre string `json:"genre,omitempty" validate:"required"`
}
