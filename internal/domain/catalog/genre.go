package catalog

import "github.com/google/uuid"

type Genre struct {
	ID       uint
	Name     string
	WebTitle string
	UUID     uuid.UUID
}
