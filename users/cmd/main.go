package main

import (
	"fmt"
	pb "github.com/blazee5/film-api/users/api/proto/v1"
	"github.com/blazee5/film-api/users/internal/config"
	"github.com/blazee5/film-api/users/internal/grpc/interceptors/auth"
	"github.com/blazee5/film-api/users/internal/grpc/interceptors/metrics"
	usergrpc "github.com/blazee5/film-api/users/internal/grpc/user"
	"github.com/blazee5/film-api/users/internal/metrics"
	"github.com/blazee5/film-api/users/internal/storage/postgres"
	sl "github.com/blazee5/film-api/users/lib/logger/slog"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()

	log := sl.SetupLogger(cfg.Env)

	log.Info("starting users service...", slog.String("env", cfg.Env))
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
		grpc.ChainUnaryInterceptor(
			auth.AuthInterceptor,
			metricsgrpc.MetricsInterceptor,
		),
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

	err = metrics.NewMetrics()
	if err != nil {
		log.Info("failed to register metrics")
	}

	go func() {
		log.Error("", metrics.ListenMetrics("0.0.0.0:2222"))
	}()

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
