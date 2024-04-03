package auth

import (
	"context"
	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	handler "go-grpc-auth/internal/handlers/auth"
	storage "go-grpc-auth/internal/storage/sqlite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

type Server struct {
	ssov1.AuthServer
	requestsHandler *handler.Handler
}

func RegisterServer(gRPC *grpc.Server, log *slog.Logger, tokenTTL time.Duration) {
	dbStorage := storage.New()
	reqHandler := handler.New(log, tokenTTL, dbStorage, dbStorage, dbStorage)
	server := Server{requestsHandler: reqHandler}
	ssov1.RegisterAuthServer(gRPC, &server)
}

func (s *Server) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}
	token, err := s.requestsHandler.Login(ctx, req.GetEmail(), req.GetPassword(), int64(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}
	userID, err := s.requestsHandler.RegisterUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (s *Server) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}
	isAdmin, err := s.requestsHandler.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
