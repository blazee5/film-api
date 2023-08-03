package postgres

import (
	pb "github.com/blazee5/film-api/api/proto"
	"github.com/blazee5/film-api/internal/models"
	"github.com/blazee5/film-api/lib/auth"
	"gorm.io/gorm"
)

const salt = "DFDjdf2434fdJFHSsdf"

func CreateUser(db *gorm.DB, in *pb.User) (id int32, err error) {
	in.Password = auth.GenerateHashPassword(in.Password, salt)

	user := &models.User{Name: in.Name, Email: in.Email, Password: in.Password}

	result := db.Create(&user)

	if result.RowsAffected == 0 {
		return 0, result.Error
	}

	return user.Id, nil
}

func GetUser(db *gorm.DB, in *pb.UserRequest) (film *pb.User, err error) {
	result := db.First(&models.User{}, in.Id)

	var userRes models.User
	result.Scan(&userRes)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.User{Id: userRes.Id, Name: userRes.Name, Email: userRes.Email}, nil
}

func UpdateUser(db *gorm.DB, in *pb.User) (film *pb.User, err error) {
	result := db.Model(&models.User{Id: in.Id}).Updates(models.User{Name: in.Name, Email: in.Email})

	var userRes models.User

	result.Scan(&userRes)

	if result.RowsAffected == 0 {
		return nil, result.Error
	}

	return &pb.User{Id: userRes.Id, Name: userRes.Name, Email: userRes.Email}, nil
}

func DeleteUser(db *gorm.DB, in *pb.UserRequest) error {
	result := db.Delete(&models.User{}, in.Id)

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
