package postgre

import (
	"mz-moon-back/internal/repository"
	pgPkg "mz-moon-back/pkg/pg"
)

type Genre struct {
	DB *pgPkg.Wrap
}

func NewGenre(pgWrap *pgPkg.Wrap) repository.IGenre {
	return &Genre{
		DB: pgWrap,
	}
}
