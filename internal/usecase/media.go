package usecase

import (
	"bytes"
	"context"
	"mz-moon-back/internal/repository"
	"os"

	"github.com/google/uuid"
)

type Media struct {
	trackRepo repository.ITrack
	coverRepo repository.ICover
}

func NewMedia(trackRepo repository.ITrack, coverRepo repository.ICover) IMedia {
	return &Media{
		trackRepo: trackRepo,
		coverRepo: coverRepo,
	}
}

var (
	ownerReadPerm os.FileMode = 0400
)

func (obj *Media) GetTrack(ctx context.Context, songUUID uuid.UUID) (*os.File, error) {
	trackData, err := obj.trackRepo.GetTrack(ctx, songUUID)
	if err != nil {
		return nil, err
	}

	trackFile, err := os.OpenFile(trackData.Path, os.O_RDONLY, ownerReadPerm)
	if err != nil {
		return nil, err
	}

	return trackFile, nil
}

func (obj *Media) GetTrackCover(ctx context.Context, songUUID uuid.UUID) (*bytes.Buffer, error) {
	coverData, err := obj.coverRepo.GetCover(ctx, songUUID)
	if err != nil {
		return nil, err
	}

	coverFile, err := os.OpenFile(coverData.Path, os.O_RDONLY, ownerReadPerm)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(coverFile)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
