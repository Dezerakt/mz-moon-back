package postgre

import (
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/repository"
	"mz-moon-back/internal/repository/models"
	pgPkg "mz-moon-back/pkg/pg"
)

type Genre struct {
	*pgPkg.Wrap
}

func NewGenre(pgWrap *pgPkg.Wrap) repository.IGenre {
	return &Genre{
		Wrap: pgWrap,
	}
}

func (obj *Genre) GetAll(ctx context.Context) ([]catalog.Genre, error) {
	var genreModels []models.Genre

	tx := obj.Db.Find(&genreModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []catalog.Genre
	for _, model := range genreModels {
		result = append(result, model.ToEntity())
	}

	return result, nil
}
