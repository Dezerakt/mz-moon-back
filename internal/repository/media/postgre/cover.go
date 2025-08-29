package postgre

import (
	"context"
	"mz-moon-back/internal/domain/media"
	"mz-moon-back/internal/repository"
	"mz-moon-back/internal/repository/models"
	pgPkg "mz-moon-back/pkg/pg"

	"github.com/google/uuid"
)

type Cover struct {
	*pgPkg.Wrap
}

func NewCover(pgWrap *pgPkg.Wrap) repository.ICover {
	return &Cover{
		Wrap: pgWrap,
	}
}

func (obj *Cover) GetCover(ctx context.Context, uuid uuid.UUID) (*media.Cover, error) {
	var cover models.Cover

	tx := obj.Db.First(&cover, "uuid = ?", uuid)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &media.Cover{
		Path: cover.Path,
	}, nil
}

func (obj *Cover) NewCover(ctx context.Context, cover *media.Cover) error {
	tx := obj.Db.Create(cover)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
