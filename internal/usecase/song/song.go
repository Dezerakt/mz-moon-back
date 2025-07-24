package song

import (
	"bytes"
	"context"
	"spotifykiller/internal/repository"
)

type Usecase struct {
	repository repository.Song
}

func NewUsecase(repository repository.Song) *Usecase {
	return &Usecase{}
}

func (obj *Usecase) Upload(ctx context.Context, songName string, buffer bytes.Buffer) {

}
