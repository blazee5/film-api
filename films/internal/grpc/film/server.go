package film_grpc

import (
	"context"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"gorm.io/gorm"
)

type Server struct {
	Db      *gorm.DB
	Service FilmService
	pb.UnimplementedFilmServiceServer
}

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name=FilmService
type FilmService interface {
	CreateFilm(db *gorm.DB, in *pb.Film) (id int64, err error)
	GetFilm(db *gorm.DB, in *pb.FilmRequest) (film *pb.Film, err error)
	UpdateFilm(db *gorm.DB, in *pb.Film) (film *pb.Film, err error)
	DeleteFilm(db *gorm.DB, in *pb.FilmRequest) error
}

func (s *Server) CreateFilm(ctx context.Context, in *pb.Film) (*pb.FilmResponse, error) {
	id, err := s.Service.CreateFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.FilmResponse{Id: id}, nil
}

func (s *Server) GetFilm(ctx context.Context, in *pb.FilmRequest) (*pb.Film, error) {
	film, err := s.Service.GetFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return film, nil
}

func (s *Server) UpdateFilm(ctx context.Context, in *pb.Film) (*pb.Film, error) {
	film, err := s.Service.UpdateFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return film, nil
}

func (s *Server) DeleteFilm(ctx context.Context, in *pb.FilmRequest) (*pb.SuccessResponse, error) {
	err := s.Service.DeleteFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Status: "success"}, nil
}
