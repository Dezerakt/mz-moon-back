package main

import (
	"accord-generator/config"
	httpDlv "accord-generator/internal/delivery/http"
	mongoRepo "accord-generator/internal/repository/mongo"
	"accord-generator/internal/usecase"
	"accord-generator/internal/usecase/mood"
	"accord-generator/internal/usecase/note"
	"accord-generator/internal/usecase/progression"
	mongoPkg "accord-generator/pkg/mongo"
	"fmt"
	"github.com/gofiber/fiber"
	"log"
)

func main() {
	// init
	app := fiber.New(&fiber.Settings{
		StrictRouting: true,
	})

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// repositories, web-api etc
	//pg := pgPkg.NewPgConnection(cfg.Postgres)
	mongoWrap := mongoPkg.NewMongoWrap(cfg.Mongo)

	// usecases
	moodUsecase := mood.NewMoodUsecase(mongoRepo.NewMoodRepo(mongoWrap))
	noteUsecase := note.NewUsecase(mongoRepo.NewNoteRepo(mongoWrap))
	progressionUsecase := progression.NewUsecase(mongoRepo.NewProgressionRepo(mongoWrap))

	uscContainer := usecase.Dependencies{ // to prevent large amount of pass arguments
		Mood:        moodUsecase,
		Note:        noteUsecase,
		Progression: progressionUsecase,
	}

	// server
	httpDlv.NewRouter(app, cfg, uscContainer)

	for _, r := range app.Routes() {
		fmt.Printf("%s\t%s\n", r.Method, r.Path)
	}

	log.Fatal(app.Listen(cfg.App.Port))
}
