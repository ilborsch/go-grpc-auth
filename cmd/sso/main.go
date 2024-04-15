package main

import (
	"go-grpc-auth/internal/app"
	"go-grpc-auth/internal/config"
	"go-grpc-auth/internal/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init config object
	cfg := config.MustLoad()

	// init logger object
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application")

	// init app and start server
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.GRPC.Timeout)
	go application.MustRun()

	// Graceful stop
	stopApp(application, log)
}

func stopApp(application *app.App, log *slog.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	stopSignal := <-stop
	log.Info("stopping application", slog.String("signal", stopSignal.String()))
	application.Stop()
}
