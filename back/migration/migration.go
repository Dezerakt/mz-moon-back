package main

import (
	"gorm.io/gorm"
	"log"
	"music-streaming/config"
	artistRepo "music-streaming/internal/repository/artist"
	genreRepo "music-streaming/internal/repository/genre"
	songRepo "music-streaming/internal/repository/song"
	pgPkg "music-streaming/pkg/pg"
)

func main() {
	log.Println("Start Migration")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Config initialized")

	pgWrap := pgPkg.NewPgConnection(cfg.Postgres)

	ArtistMigrate(pgWrap)
	GenreMigration(pgWrap)
	SongMigration(pgWrap)

	log.Println("Migration ended")
}

func SongMigration(wrap *pgPkg.PgWrap) {
	log.Println("Start `Song` model migration")
	defer log.Println("`Song` model migration ended")

	model := &songRepo.Song{}

	err := wrap.Db.Migrator().DropTable(model)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wrap.Db.Migrator().AutoMigrate(model)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GenreMigration(wrap *pgPkg.PgWrap) {
	log.Println("Start `Genre` model migration")
	defer log.Println("`Genre` model migration ended")

	model := &genreRepo.Genre{}

	err := wrap.Db.Migrator().DropTable(model)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wrap.Db.Migrator().AutoMigrate(model)
	if err != nil {
		log.Fatal(err)
		return
	}

	genres := []*genreRepo.Genre{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "rap",
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
			Name: "break beat",
		},
	}

	wrap.Db.Create(genres)
}

func ArtistMigrate(wrap *pgPkg.PgWrap) {
	log.Println("Start `Artist` model migration")
	defer log.Println("`Artist` model migration ended")

	model := &artistRepo.Artist{}

	err := wrap.Db.Migrator().DropTable(model)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = wrap.Db.Migrator().AutoMigrate(model)
	if err != nil {
		log.Fatal(err)
		return
	}

	artist := []*artistRepo.Artist{
		{
			Model: gorm.Model{
				ID: 1,
			},
			AristName: "lil peep",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			AristName: "hazzequill",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			AristName: "Nirvana",
		},
	}

	wrap.Db.Create(artist)
}
