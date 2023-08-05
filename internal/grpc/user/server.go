package user_grpc

import (
	"context"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"github.com/blazee5/film-api/internal/storage/postgres"
	"github.com/blazee5/film-api/lib/auth"
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
	pb.UnimplementedUserServiceServer
}

const (
	salt = "DFDjdf2434fdJFHSsdf"
)

func (s *Server) SignUp(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	id, err := postgres.CreateUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.User) (*pb.Token, error) {
	user, err := postgres.ValidateUser(s.Db, in.Email, auth.GenerateHashPassword(in.Password, salt))

	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(s.Db, user.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func (s *Server) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.User, error) {
	user, err := postgres.GetUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := postgres.UpdateUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Server) DeleteUser(ctx context.Context, in *pb.UserRequest) (*pb.SuccessResponse, error) {
	err := postgres.DeleteUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Status: "success"}, nil
}
