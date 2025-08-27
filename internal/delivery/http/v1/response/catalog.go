package response

import "github.com/google/uuid"

type GetAllSongs struct {
	Song  string    `json:"name,omitempty" validate:"required"`
	Arist string    `json:"artist,omitempty" validate:"required"`
	Genre string    `json:"genre,omitempty" validate:"required"`
	UUID  uuid.UUID `json:"uuid,omitempty" validate:"required"`
}

type GetAllGenres struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}
