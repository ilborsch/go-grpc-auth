package models

import (
	"context"
)

type User struct {
	ID             int64
	Email          string
	PasswordHashed []byte
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passwordHash []byte,
	) (int64, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
