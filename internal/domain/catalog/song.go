package catalog

import "github.com/google/uuid"

type Song struct {
	UUID     uuid.UUID
	SongName string
	Artist   string
	Genre    string
}
