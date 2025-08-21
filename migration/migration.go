package main

import (
	"fmt"
	"log"
	"mz-moon-back/config"
	"mz-moon-back/internal/repository/models"

	"github.com/google/uuid"
	"gorm.io/gorm"

	pgPkg "mz-moon-back/pkg/pg"
)

var (
	mockUUID = "7f1bddad-08fc-4eac-b550-d3aee7a4ecb8"
)

func main() {
	log.Println("Start Migration")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config initialized")

	pgWrap := pgPkg.NewPgConnection(cfg.Postgres)

	CatalogMigration(pgWrap)
	MediaMigration(pgWrap)

	log.Println("Migration ended")
}

func CatalogMigration(wrap *pgPkg.Wrap) {
	log.Println("Start `Catalog` model migration")
	defer log.Println("`Catalog` model migration ended")

	songModel := &models.Song{}
	artistModel := &models.Artist{}
	genreModel := &models.Genre{}

	err := wrap.Db.Migrator().DropTable(songModel, artistModel, genreModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wrap.Db.Migrator().AutoMigrate(songModel, artistModel, genreModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	genre := []models.Genre{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "trap",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "grange",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "breakcore",
		},
	}

	artist := []models.Artist{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Nirvana",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "hazzequill",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "lil peep",
		},
	}

	parsedUUID, _ := uuid.Parse(mockUUID)
	songs := []models.Song{
		{
			Model: gorm.Model{
				ID: 1,
			},
			ArtistID: 3,
			GenreID:  1,
			Name:     "benz truck",
			UUID:     parsedUUID,
		},
	}

	wrap.Db.Create(&genre)
	wrap.Db.Create(&artist)
	wrap.Db.Create(&songs)
}

func MediaMigration(wrap *pgPkg.Wrap) {
	log.Println("Start `Media` model migration")
	defer log.Println("`Media` model migration ended")

	trackModel := &models.Track{}
	coverModel := &models.Cover{}

	err := wrap.Db.Migrator().DropTable(trackModel, coverModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wrap.Db.Migrator().AutoMigrate(trackModel, coverModel)
	if err != nil {
		log.Fatal(err)
		return
	}

	parsedUUID, _ := uuid.Parse(mockUUID)
	tracks := []models.Track{
		{
			Model: gorm.Model{
				ID: 1,
			},
			UUID: parsedUUID,
			Path: fmt.Sprintf("/storage/song/%s.mp3", mockUUID),
		},
	}

	covers := []models.Cover{
		{
			Model: gorm.Model{
				ID: 1,
			},
			UUID: parsedUUID,
			Path: fmt.Sprintf("/storage/cover/%s.jpeg", mockUUID),
		},
	}

	wrap.Db.Create(&tracks)
	wrap.Db.Create(&covers)
}
