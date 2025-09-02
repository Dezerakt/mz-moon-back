package catalog

import "github.com/google/uuid"

type Song struct {
	SongName string
	Artist   string
	Genre    string
	UUID     uuid.UUID
}
