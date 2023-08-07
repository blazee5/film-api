package main

import (
	"context"
	"fmt"
	pb "github.com/blazee5/film-api/api/proto/v1"
	"github.com/blazee5/film-api/internal/config"
	usergrpc "github.com/blazee5/film-api/internal/grpc/user"
	"github.com/blazee5/film-api/internal/grpc/user/interceptors/auth"
	"github.com/blazee5/film-api/internal/rabbitmq"
	"github.com/blazee5/film-api/internal/storage/postgres"
	sl "github.com/blazee5/film-api/lib/logger/slog"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting user server...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Info("error while connecting database", sl.Err(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Info("failed to listen:", err)
	}
	s := grpc.NewServer(
		withServerUnaryInterceptor(),
	)

	pb.RegisterUserServiceServer(s, &usergrpc.Server{
		Db:      db.Db,
		Service: &db,
	})

	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Info("failed to serve:", err)
		}
	}()

	conn, err := rabbitmq.NewRabbitMQConn(cfg)
	if err != nil {
		log.Info("failed to connect rabbitmq:", err)
	}

	ch, err := rabbitmq.NewChannelConn(conn, log)

	defer ch.Close()

	q, err := rabbitmq.NewQueueConn(ch, log)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	err = rabbitmq.SendMessage(ctx, "Hello", ch, q)

	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.GracefulStop()

	dbCon, err := db.Db.DB()
	if err != nil {
		log.Info("error in db:", sl.Err(err))
	}

	err = dbCon.Close()
	if err != nil {
		log.Info("error while close db:", sl.Err(err))
	}

}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(auth.ServerInterceptor)
}
