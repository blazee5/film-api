package user_grpc

import (
	"context"
	"errors"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"github.com/blazee5/film-api/users/internal/models"
	"github.com/blazee5/film-api/users/lib/auth"
	"gorm.io/gorm"
)

type Server struct {
	Db      *gorm.DB
	Service UserService
	pb.UnimplementedUserServiceServer
}

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name=UserService
type UserService interface {
	CreateUser(db *gorm.DB, in *pb.User) (int64, error)
	ValidateUser(db *gorm.DB, email, password string) (*models.User, error)
	GetUser(db *gorm.DB, in *pb.UserRequest) (*pb.UserInfo, error)
	UpdateUser(db *gorm.DB, user *pb.User) (*pb.User, error)
	DeleteUser(db *gorm.DB, in *pb.UserRequest) error
}

func (s *Server) SignUp(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	id, err := s.Service.CreateUser(s.Db, in)

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, errors.New("email already use")
	}

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.User) (*pb.Token, error) {
	user, err := s.Service.ValidateUser(s.Db, in.Email, auth.GenerateHashPassword(in.Password))

	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(s.Db, user.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserInfo, error) {
	if ctx.Value("user_id") != int(in.Id) {
		return nil, errors.New("you are not this users")
	}
	user, err := s.Service.GetUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := s.Service.UpdateUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.UserRequest) (*pb.SuccessResponse, error) {
	err := s.Service.DeleteUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Status: "success"}, nil
}
