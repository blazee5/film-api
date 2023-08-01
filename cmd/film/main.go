package main

import (
	"fmt"
	pb "github.com/blazee5/film-api/api/proto"
	"github.com/blazee5/film-api/internal/config"
	filmgrpc "github.com/blazee5/film-api/internal/grpc/film"
	"github.com/blazee5/film-api/internal/storage/postgres"
	sl "github.com/blazee5/film-api/lib/logger/slog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.LoadConfig()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting film server...", slog.String("env", cfg.Env))
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
	pb.RegisterFilmServiceServer(s, &filmgrpc.Server{Db: db})
	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))
	if err := s.Serve(lis); err != nil {
		log.Info("failed to serve: %v", err)
	}
}
