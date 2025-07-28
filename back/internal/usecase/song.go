package usecase

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"music-streaming/internal/domain/song"
	"music-streaming/internal/repository"
	"os"
	"path/filepath"
	"strings"
)

type Song struct {
	repository repository.ISong
}

func NewSongUsecase(repository repository.ISong) ISong {
	return &Song{
		repository: repository,
	}
}

func (obj *Song) Upload(ctx context.Context, songEntity *song.Entity, buffer *bytes.Buffer) (uint, error) {
	songFileName := fmt.Sprintf("%s.mp3", formatSongName(songEntity.SongName))
	songPath := filepath.Join("storage", "song", songFileName)

	log.Printf("Размер буфера: %d байт", buffer.Len())
	log.Println(songPath)

	if err := os.MkdirAll(filepath.Dir(songPath), os.ModePerm); err != nil {
		return 0, err
	}

	if err := os.WriteFile(songPath, buffer.Bytes(), 0644); err != nil {
		return 0, err
	}

	insertedID, err := obj.repository.CreateNewMeta(ctx, songEntity, songPath)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func formatSongName(rawName string) string {
	formattedString := strings.ReplaceAll(rawName, " ", "_")
	formattedString = strings.ToLower(formattedString)

	return formattedString
}
