package catalog

import "github.com/google/uuid"

type Genre struct {
	ID   uint
	Name string
	UUID uuid.UUID
}
