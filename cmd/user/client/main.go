package main

import (
	"fmt"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"github.com/blazee5/film-api/internal/config"
	sl "github.com/blazee5/film-api/lib/logger/slog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.LoadConfig()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting user service client...")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Info("err:", sl.Err(err))
	}

	defer conn.Close()

	_ = pb.NewUserServiceClient(conn)
}
