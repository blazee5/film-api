package postgres

import (
	"context"
	pb "github.com/blazee5/film-api/films/api/proto/v1"
	"github.com/blazee5/film-api/films/internal/models"
	"gorm.io/gorm"
)

func (p *Postgres) CreateFilm(ctx context.Context, db *gorm.DB, in *pb.Film) (id int64, err error) {
	film := &models.Film{Title: in.Title, Description: in.Description, Genre: in.Genre}

	result := db.Create(&film).WithContext(ctx)

	if result.RowsAffected == 0 {
		return 0, result.Error
	}

	return film.Id, nil
}

func (p *Postgres) GetFilm(ctx context.Context, db *gorm.DB, in *pb.FilmRequest) (film *pb.Film, err error) {
	result := db.First(&models.Film{}, in.Id).WithContext(ctx)

	var filmRes models.Film
	result.Scan(&filmRes)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.Film{Id: filmRes.Id, Title: filmRes.Title, Description: filmRes.Description, Genre: filmRes.Genre}, nil
}

func (p *Postgres) UpdateFilm(ctx context.Context, db *gorm.DB, in *pb.Film) (film *pb.Film, err error) {
	result := db.Model(&models.Film{Id: in.Id}).Updates(models.Film{Title: in.Title, Description: in.Description, Genre: in.Genre}).WithContext(ctx)

	var filmRes models.Film

	result.Scan(&filmRes)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.Film{Id: filmRes.Id, Title: filmRes.Title, Description: filmRes.Description, Genre: filmRes.Genre}, nil
}

func (p *Postgres) DeleteFilm(ctx context.Context, db *gorm.DB, in *pb.FilmRequest) error {
	result := db.Delete(&models.Film{}, in.Id).WithContext(ctx)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
