package postgres

import (
	"errors"
	pb "github.com/blazee5/film-api/users/api/proto/v1"
	"github.com/blazee5/film-api/users/internal/models"
	"github.com/blazee5/film-api/users/lib/auth"
	"gorm.io/gorm"
)

func (p *Postgres) CreateUser(db *gorm.DB, in *pb.User) (id int64, err error) {
	in.Password = auth.GenerateHashPassword(in.Password)

	user := &models.User{Name: in.Name, Email: in.Email, Password: in.Password}

	result := db.Create(&user)

	if result.RowsAffected == 0 {
		return 0, result.Error
	}

	return user.Id, nil
}

func (p *Postgres) GetUser(db *gorm.DB, in *pb.UserRequest) (film *pb.UserInfo, err error) {
	result := db.First(&models.User{}, in.Id)

	var userRes models.User
	result.Scan(&userRes)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.UserInfo{Id: userRes.Id, Name: userRes.Name, Email: userRes.Email}, nil
}

func (p *Postgres) UpdateUser(db *gorm.DB, in *pb.User) (film *pb.User, err error) {
	result := db.Model(&models.User{Id: in.Id}).Updates(models.User{Name: in.Name, Email: in.Email})

	var userRes models.User

	result.Scan(&userRes)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.User{Id: userRes.Id, Name: userRes.Name, Email: userRes.Email}, nil
}

func (p *Postgres) DeleteUser(db *gorm.DB, in *pb.UserRequest) error {
	result := db.Delete(&models.User{}, in.Id)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (p *Postgres) ValidateUser(db *gorm.DB, email, password string) (*models.User, error) {
	var user *models.User

	result := db.Where("email = ? AND password = ?", email, password).Find(&user)

	if result.RowsAffected == 0 {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
