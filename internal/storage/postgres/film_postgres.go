package postgres

import (
	pb "github.com/blazee5/film-api/internal/film/api/film_v1"
	"github.com/blazee5/film-api/internal/models"
	"gorm.io/gorm"
)

func CreateFilm(db *gorm.DB, in *pb.Film) (id int32, err error) {
	film := &models.Film{Title: in.Title, Description: in.Description, Genre: in.Genre}

	result := db.Create(&film)

	if result.RowsAffected == 0 {
		return 0, result.Error
	}

	return film.Id, nil
}

func GetFilm(db *gorm.DB, in *pb.FilmRequest) (film *models.Film, err error) {
	result := db.First(&models.Film{}, in.Id)

	var filmRes models.Film
	result.Scan(&filmRes)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &models.Film{Id: filmRes.Id, Title: filmRes.Title, Description: filmRes.Description, Genre: filmRes.Genre}, nil
}

func UpdateFilm(db *gorm.DB, in *pb.Film) (film *models.Film, err error) {
	result := db.Model(&models.Film{Id: in.Id}).Updates(models.Film{Title: in.Title, Description: in.Description, Genre: in.Genre})

	var filmRes models.Film

	result.Scan(&filmRes)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &models.Film{Id: filmRes.Id, Title: filmRes.Title, Description: filmRes.Description, Genre: filmRes.Genre}, nil
}

func DeleteFilm(db *gorm.DB, in *pb.FilmRequest) error {
	result := db.Delete(&models.Film{}, in.Id)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}