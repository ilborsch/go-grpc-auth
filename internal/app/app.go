package app

import (
	"fmt"
	authgrpc "go-grpc-auth/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type App struct {
	GRPCServer *grpc.Server
	log        *slog.Logger
	port       int
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.RegisterServer(gRPCServer)
	return &App{
		log:        log,
		port:       grpcPort,
		GRPCServer: gRPCServer,
	}
}

func (app *App) runServer() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", app.port))
	if err != nil {
		return fmt.Errorf("grpcapp.Run: %v", err.Error())
	}
	app.log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := app.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("grpcapp.Run: %v", err.Error())
	}
	return nil
}

func (app *App) MustRun() {
	if err := app.runServer(); err != nil {
		panic("App.MustRun error: " + err.Error())
	}
}

func (app *App) Stop() {
	app.log.Info("stopping gRPC server", slog.Int("port", app.port))
	app.GRPCServer.GracefulStop()
}
