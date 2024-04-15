package handler

import (
	"context"
	tkn "go-grpc-auth/internal/jwt"
	"go-grpc-auth/internal/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

const logFrom = "handlers/auth/Handler"

type Handler struct {
	log      *slog.Logger
	tokenTTL time.Duration
	models.UserSaver
	models.UserProvider
	models.AppProvider
}

func (h *Handler) Login(ctx context.Context, email, password string, appID int64) (string, error) {
	log := h.log.With(
		slog.String("from", logFrom+".Login"),
		slog.String("email", email),
	)
	log.Info("logging user in")
	user, err := h.UserProvider.User(ctx, email)
	if err != nil {
		log.Error("error retrieving user from storage " + err.Error())
		return "", status.Error(codes.InvalidArgument, "invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword(user.PasswordHashed, []byte(password)); err != nil {
		log.Error("password and hashed passwords don't match " + err.Error())
		return "", status.Error(codes.InvalidArgument, "wrong credentials")
	}
	app, err := h.AppProvider.App(ctx, appID)
	if err != nil {
		log.Error("error retrieving app from storage " + err.Error())
		return "", status.Error(codes.InvalidArgument, "invalid app_id")
	}
	token, err := tkn.NewToken(user, app, h.tokenTTL)
	if err != nil {
		log.Error("error creating jwt token " + err.Error())
		return "", status.Error(codes.Internal, "internal error")
	}
	return token, nil
}

func (h *Handler) RegisterUser(ctx context.Context, email, password string) (int64, error) {
	log := h.log.With(
		slog.String("from", logFrom+".RegisterUser"),
		slog.String("email", email),
	)
	log.Info("registering user")
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
	log := h.log.With(
		slog.String("from", logFrom+".IsAdmin"),
		slog.Int("uID", int(userID)),
	)
	log.Info("getting user permissions")
	isAdmin, err := h.IsAdmin(ctx, userID)
	if err != nil {
		log.Error("error retrieving user permissions " + err.Error())
		return false, status.Error(codes.Internal, "internal error")
	}
	return isAdmin, nil
}

func New(
	log *slog.Logger,
	tokenTTL time.Duration,
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
