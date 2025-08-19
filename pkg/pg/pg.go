package pgPkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mz-moon-back/config"
)

type PgWrap struct {
	Db *gorm.DB
}

func NewPgConnection(dbCfg config.Postgres) *PgWrap {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Db, dbCfg.Port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("Connected to PostgreSQL successfully")
	return &PgWrap{
		Db: db,
	}
}
