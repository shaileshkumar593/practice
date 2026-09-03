package main

import (
	"context"
	apporder "grpc-order-service/internal/application/order"
	"grpc-order-service/internal/config"
	grpcserver "grpc-order-service/internal/grpc"
	"grpc-order-service/internal/infrastructure/postgres"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.Load()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	db, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ping, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.Ping(ping); err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewOrderRepository(db)
	svc := apporder.New(repo)
	handler := grpcserver.NewOrderHandler(svc)
	server := grpcserver.NewServer(handler, logger)
	go func() {
		if err := grpcserver.ListenAndServe(server, cfg.Port); err != nil {
			logger.Error("grpc server failed", "error", err)
			stop()
		}
	}()
	<-ctx.Done()
	logger.Info("shutting down")
	server.GracefulStop()
}
