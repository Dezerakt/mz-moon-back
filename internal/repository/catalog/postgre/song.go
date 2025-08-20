package postgre

import (
	"mz-moon-back/internal/repository"
	pgPkg "mz-moon-back/pkg/pg"
)

type Song struct {
	DB *pgPkg.Wrap
}

func NewSong(pgWrap *pgPkg.Wrap) repository.ISong {
	return &Song{
		DB: pgWrap,
	}
}
