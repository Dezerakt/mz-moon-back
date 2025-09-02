package catalog

import "github.com/google/uuid"

type Artist struct {
	ID   int
	Name string
	UUID uuid.UUID
}
