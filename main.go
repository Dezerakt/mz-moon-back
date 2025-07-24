package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"log"
	"spotifykiller/config"
	httpDlv "spotifykiller/internal/delivery/http"
	pgRepo "spotifykiller/internal/repository/pg"
	"spotifykiller/internal/usecase/song"
	pgPkg "spotifykiller/pkg/pg"
)

func main() {
	// init
	app := fiber.New()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// repositories, web-api etc
	pg := pgPkg.NewPgConnection(cfg.Postgres)

	// usecases
	songUsecase := song.NewUsecase(pgRepo.NewSongRepo(pg))

	// server
	httpDlv.NewRouter(app, cfg, songUsecase)

	for _, r := range app.Routes() {
		fmt.Printf("%s\t%s\n", r.Method, r.Path)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
