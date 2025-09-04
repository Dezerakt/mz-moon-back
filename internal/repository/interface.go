package repository

import (
	"bytes"
	"context"
	"mz-moon-back/internal/domain/catalog"
	"mz-moon-back/internal/domain/media"

	"github.com/google/uuid"
)

type (
	ISong interface {
		GetAll(ctx context.Context) ([]catalog.Song, error)
	}

	IGenre interface {
		GetAll(ctx context.Context) ([]catalog.Genre, error)
		NewGenre(ctx context.Context, genre *catalog.Genre) error
	}

	IArtist interface {
	}

	ITrack interface {
		GetTrack(ctx context.Context, uuid uuid.UUID) (*media.Track, error)
	}

	ICover interface {
		UpsertCover(ctx context.Context, cover *media.Cover) error
		GetCover(ctx context.Context, uuid uuid.UUID) (*media.Cover, error)
	}

	IFileManager interface {
		StoreCover(ctx context.Context, uuid uuid.UUID, contentType media.ContentType, fileContent *bytes.Buffer) error
	}
)
