package film

import (
	"context"
	pb "github.com/blazee5/film-api/internal/film/api/film_v1"
	"github.com/blazee5/film-api/internal/storage/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
	pb.UnimplementedFilmServiceServer
}

func (s *Server) CreateFilm(ctx context.Context, in *pb.Film) (*pb.FilmResponse, error) {
	id, err := postgres.CreateFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.FilmResponse{Id: id}, nil
}

func (s *Server) GetFilm(ctx context.Context, in *pb.FilmRequest) (*pb.Film, error) {
	film, err := postgres.GetFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return film, nil
}

func (s *Server) UpdateFilm(ctx context.Context, in *pb.Film) (*pb.Film, error) {
	film, err := postgres.UpdateFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return film, nil
}

func (s *Server) DeleteFilm(ctx context.Context, in *pb.FilmRequest) (*pb.SuccessResponse, error) {
	err := postgres.DeleteFilm(s.Db, in)

	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Status: "success"}, nil
}
