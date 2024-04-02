package handler

import (
	"context"
	"go-grpc-auth/internal/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

const logFrom = "handlers/auth/Handler"

type Handler struct {
	log *slog.Logger
	models.UserSaver
	models.UserProvider
	models.AppProvider
}

func (h *Handler) Login(ctx context.Context, email, password string, appID int32) (string, error) {
	log := h.log.With(
		slog.String("from", logFrom+".Login"),
		slog.String("email", email),
	)
	user, err := h.UserProvider.User(ctx, email)
	if err != nil {
		log.Error("error retrieving user from storage " + err.Error())
		return "", status.Error(codes.Internal, "internal error")
	}
	if err := bcrypt.CompareHashAndPassword(user.PasswordHashed, []byte(password)); err != nil {
		log.Error("password and hashed passwords don't match " + err.Error())
		return "", status.Error(codes.InvalidArgument, "wrong credentials")
	}
	// TODO: 1:56:39
	panic("implement me")
}

func (h *Handler) RegisterUser(ctx context.Context, email, password string) (int64, error) {
	log := h.log.With(
		slog.String("from", logFrom+".RegisterUser"),
		slog.String("email", email),
	)
	passHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error while hashing password " + err.Error())
		return 0, status.Error(codes.Internal, "internal error")
	}
	id, err := h.UserSaver.SaveUser(ctx, email, passHashed)
	if err != nil {
		log.Error("error creating user object " + err.Error())
		return 0, status.Error(codes.Internal, "internal error")
	}
	log.Info("user registered")
	return id, nil
}

func (h *Handler) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("implement me")
}

func New(
	log *slog.Logger,
	userSaver models.UserSaver,
	userProvider models.UserProvider,
	appProvider models.AppProvider,
) *Handler {
	return &Handler{
		log:          log,
		UserSaver:    userSaver,
		UserProvider: userProvider,
		AppProvider:  appProvider,
	}
}
