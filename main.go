package main

import (
	"fmt"
	"log"
	"mz-moon-back/config"
	httpDlv "mz-moon-back/internal/delivery/http"
	"mz-moon-back/internal/repository/catalog/postgre"
	"mz-moon-back/internal/usecase"
	pgPkg "mz-moon-back/pkg/pg"

	"github.com/gofiber/fiber"
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

	// connection objects
	pgConnection := pgPkg.NewPgConnection(cfg.Postgres)

	// repositories, web-api etc
	songRepo := postgre.NewSong(pgConnection)
	artistRepo := postgre.NewArtist(pgConnection)
	genreRepo := postgre.NewGenre(pgConnection)

	// usecases
	catalogUsecase := usecase.NewCatalog(songRepo, artistRepo, genreRepo)

	// server
	httpDlv.NewRouter(app, cfg, catalogUsecase)

	for _, r := range app.Routes() {
		fmt.Printf("%s\t%s\n", r.Method, r.Path)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
