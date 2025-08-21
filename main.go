package main

import (
	"fmt"
	"log"
	"mz-moon-back/config"
	httpDlv "mz-moon-back/internal/delivery/http"
	catalogPg "mz-moon-back/internal/repository/catalog/postgre"
	mediaPg "mz-moon-back/internal/repository/media/postgre"
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
	songRepo := catalogPg.NewSong(pgConnection)
	artistRepo := catalogPg.NewArtist(pgConnection)
	genreRepo := catalogPg.NewGenre(pgConnection)
	trackRepo := mediaPg.NewTrack(pgConnection)
	coverRepo := mediaPg.NewCover(pgConnection)

	// usecases
	catalogUsecase := usecase.NewCatalog(songRepo, artistRepo, genreRepo)
	mediaUsecase := usecase.NewMedia(trackRepo, coverRepo)

	// server
	httpDlv.NewRouter(app, cfg, catalogUsecase, mediaUsecase)

	for _, r := range app.Routes() {
		fmt.Printf("%s\t%s\n", r.Method, r.Path)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
