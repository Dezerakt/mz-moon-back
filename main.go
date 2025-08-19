package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"log"
	"music-streaming/config"
	httpDlv "music-streaming/internal/delivery/http"
	artistRepo "music-streaming/internal/repository/artist"
	songRepo "music-streaming/internal/repository/song"
	"music-streaming/internal/usecase"
	pgPkg "music-streaming/pkg/pg"
)

func main() {
	// init
	app := fiber.New(&fiber.Settings{
		BodyLimit: 10 * 1024 * 1024,
	})

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// repositories, web-api etc
	pg := pgPkg.NewPgConnection(cfg.Postgres)

	// usecases
	songUsecase := usecase.NewSongUsecase(songRepo.NewSongRepo(pg))
	artistUsecase := usecase.NewArtistUsecase(artistRepo.NewArtistRepo(pg))

	// server
	httpDlv.NewRouter(app, cfg, songUsecase, artistUsecase)

	for _, r := range app.Routes() {
		fmt.Printf("%s\t%s\n", r.Method, r.Path)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
