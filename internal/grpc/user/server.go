package user_grpc

import (
	"context"
	pb "github.com/blazee5/film-api/api/proto"
	"github.com/blazee5/film-api/internal/storage/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.UserResponse, error) {
	id, err := postgres.CreateUser(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id}, nil
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
