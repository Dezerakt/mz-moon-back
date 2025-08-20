package main

import (
	"log"
	"mz-moon-back/config"
	"mz-moon-back/internal/repository/models"

	"gorm.io/gorm"

	pgPkg "mz-moon-back/pkg/pg"
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

	log.Println("Migration ended")
}

func CatalogMigration(wrap *pgPkg.Wrap) {
	log.Println("Start `Song` model migration")
	defer log.Println("`Song` model migration ended")

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

	songs := []models.Song{
		{
			Model: gorm.Model{
				ID: 1,
			},
			ArtistID: 3,
			GenreID:  1,
			Name:     "benz truck",
			//FilePath: "/storage/song/benz_truck.mp3",
		},
	}

	wrap.Db.Create(&genre)
	wrap.Db.Create(&artist)
	wrap.Db.Create(&songs)
}
