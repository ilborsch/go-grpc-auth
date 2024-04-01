package auth

import "context"

type Auth interface {
	Login(ctx context.Context, email, password string, appID int32) (string, error)
	RegisterUser(ctx context.Context, email, password string) (int64, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
