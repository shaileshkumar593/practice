package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"boilerplate/internal/config"
	"boilerplate/internal/database"
	"boilerplate/internal/endpoint"
	appmiddleware "boilerplate/internal/middleware"
	"boilerplate/internal/repository"
	"boilerplate/internal/service"
	httptransport "boilerplate/internal/transport/http"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {

	// --------------------------------------------------
	// Configuration
	// --------------------------------------------------

	cfg := config.Load()

	// --------------------------------------------------
	// Logger
	// --------------------------------------------------

	logger := kitlog.NewLogfmtLogger(os.Stderr)

	logger = level.NewFilter(
		logger,
		level.AllowInfo(),
	)

	_ = level.Info(logger).Log(
		"msg", "starting application",
		"name", cfg.AppName,
		"port", cfg.HTTPPort,
	)

	// --------------------------------------------------
	// Database
	// --------------------------------------------------

	db, err := database.NewPostgres(cfg)
	if err != nil {
		_ = level.Error(logger).Log(
			"msg", "database connection failed",
			"error", err,
		)

		os.Exit(1)
	}

	defer db.Close()

	// --------------------------------------------------
	// Repository
	// --------------------------------------------------

	Repo := repository.NewRepository(db)

	// --------------------------------------------------
	// Service
	// --------------------------------------------------

	userService := service.NewService(Repo)

	// --------------------------------------------------
	// Endpoint
	// --------------------------------------------------

	userEndpoint := endpoint.MakeEndpoints(userService)

	// --------------------------------------------------
	// HTTP Transport
	// --------------------------------------------------

	handler := httptransport.NewUserHandler(userEndpoint)

	// --------------------------------------------------
	// Middleware
	// --------------------------------------------------

	handler = appmiddleware.Logging(
		handler,
		logger,
	)

	// --------------------------------------------------
	// HTTP Server
	// --------------------------------------------------

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,

		ReadHeaderTimeout: 5 * time.Second,

		ReadTimeout: 10 * time.Second,

		WriteTimeout: 10 * time.Second,

		IdleTimeout: 60 * time.Second,
	}

	// --------------------------------------------------
	// Start HTTP Server
	// --------------------------------------------------

	go func() {

		_ = level.Info(logger).Log(
			"msg", "HTTP server listening",
			"port", cfg.HTTPPort,
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			_ = level.Error(logger).Log(
				"msg", "HTTP server stopped",
				"error", err,
			)

			os.Exit(1)
		}
	}()

	// --------------------------------------------------
	// Graceful Shutdown
	// --------------------------------------------------

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	_ = level.Info(logger).Log(
		"msg", "shutdown signal received",
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		_ = level.Error(logger).Log(
			"msg", "graceful shutdown failed",
			"error", err,
		)

		return
	}

	_ = level.Info(logger).Log(
		"msg", "server shutdown completed",
	)
}
