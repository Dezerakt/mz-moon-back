package fileManager

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mz-moon-back/internal/domain/media"
	"mz-moon-back/internal/repository"
	"os"

	"github.com/google/uuid"
)

type FileManager struct {
}

func NewFileManager() repository.IFileManager {
	return &FileManager{}
}

const (
	ownerReadPerm os.FileMode = 0400
	coverStorage              = "/storage/cover"
	songStorage               = "/storage/song"
)

func (obj *FileManager) StoreCover(ctx context.Context, uuid uuid.UUID, contentType media.ContentType, fileContent *bytes.Buffer) error {
	var (
		fileName string
	)

	switch contentType {
	case media.GenreCover:
		fileName = fmt.Sprintf("%s/genre_%s.jpeg", coverStorage, uuid.String())
	default:
		return errors.New("invalid file content type")
	}

	err := os.WriteFile(fileName, fileContent.Bytes(), ownerReadPerm)
	if err != nil {
		return err
	}

	return nil
}
