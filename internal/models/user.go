package models

import (
	"context"
)

type User struct {
	ID             int
	Email          string
	PasswordHashed []byte
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passwordHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
