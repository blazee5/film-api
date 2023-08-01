package main

import (
	"context"
	"fmt"
	"github.com/blazee5/film-api/internal/config"
	pb "github.com/blazee5/film-api/internal/film/api/film_v1"
	"github.com/blazee5/film-api/internal/storage/postgres"
	sl "github.com/blazee5/film-api/lib/logger/slog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type server struct {
	db *gorm.DB
	pb.UnimplementedFilmServiceServer
}

func (s *server) CreateFilm(ctx context.Context, in *pb.Film) (*pb.FilmResponse, error) {
	id, err := postgres.CreateFilm(s.db, in)

	if err != nil {
		return nil, err
	}

	return &pb.FilmResponse{Id: id}, nil
}

func (s *server) GetFilm(ctx context.Context, in *pb.FilmRequest) (*pb.Film, error) {
	film, err := postgres.GetFilm(s.db, in)

	if err != nil {
		return nil, err
	}

	return &pb.Film{Title: film.Title, Description: film.Description, Genre: film.Genre}, nil
}

func (s *server) UpdateFilm(ctx context.Context, in *pb.Film) (*pb.Film, error) {
	film, err := postgres.UpdateFilm(s.db, in)

	if err != nil {
		return nil, err
	}

	return &pb.Film{Title: film.Title, Description: film.Description, Genre: film.Genre}, nil
}

func (s *server) DeleteFilm(ctx context.Context, in *pb.FilmRequest) (*pb.SuccessResponse, error) {
	err := postgres.DeleteFilm(s.db, in)

	if err != nil {
		return nil, err
	}

	return &pb.SuccessResponse{Status: "success"}, nil
}

func main() {
	cfg := config.LoadConfig()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting ecommerce server...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Info("error while connecting database", sl.Err(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Info("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFilmServiceServer(s, &server{db: db})
	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))
	if err := s.Serve(lis); err != nil {
		log.Info("failed to serve: %v", err)
	}
}
