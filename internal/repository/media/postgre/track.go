package postgre

import (
	"context"
	"mz-moon-back/internal/domain/media"
	"mz-moon-back/internal/repository"
	"mz-moon-back/internal/repository/models"
	pgPkg "mz-moon-back/pkg/pg"

	"github.com/google/uuid"
)

type Track struct {
	*pgPkg.Wrap
}

func NewTrack(pgWrap *pgPkg.Wrap) repository.ITrack {
	return &Track{
		Wrap: pgWrap,
	}
}

func (obj *Track) GetTrack(ctx context.Context, uuid uuid.UUID) (*media.Track, error) {
	var track models.Track

	tx := obj.Db.First(&track, "uuid = ?", uuid)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &media.Track{
		UUID: track.UUID,
	}, nil
}
