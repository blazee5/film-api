package postgres

import (
	"fmt"
	"github.com/blazee5/film-api/films/internal/config"
	"github.com/blazee5/film-api/films/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Film{})

	if err != nil {
		panic(err)
	}
	return &Postgres{Db: db}, nil
}
