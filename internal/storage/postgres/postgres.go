package postgres

import (
	"fmt"
	"github.com/blazee5/film-api/internal/config"
	"github.com/blazee5/film-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres(cfg *config.Config) (Postgres, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return Postgres{}, err
	}

	err = db.AutoMigrate(&models.Film{})
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return Postgres{}, err
	}
	return Postgres{Db: db}, nil
}
