package postgre

import (
	"mz-moon-back/internal/repository"
	pgPkg "mz-moon-back/pkg/pg"
)

type Artist struct {
	DB *pgPkg.Wrap
}

func NewArtist(pgWrap *pgPkg.Wrap) repository.IArtist {
	return &Artist{
		DB: pgWrap,
	}
}
