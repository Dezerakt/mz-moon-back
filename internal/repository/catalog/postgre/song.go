package postgre

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/repository"
	"mz-moon-back/internal/repository/models"
	pgPkg "mz-moon-back/pkg/pg"
)

type Song struct {
	*pgPkg.Wrap
}

func NewSong(pgWrap *pgPkg.Wrap) repository.ISong {
	return &Song{
		Wrap: pgWrap,
	}
}

func (obj *Song) GetAll(ctx context.Context) ([]catalog.Song, error) {
	var songModels []models.Song

	tx := obj.Db.
		Joins("Artist").
		Joins("Genre").
		Find(&songModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []catalog.Song
	for _, model := range songModels {
		result = append(result, model.ToEntity())
	}

	return result, nil
}
