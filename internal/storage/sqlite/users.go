package storage

import (
	"context"
	"go-grpc-auth/internal/models"
)

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const query = `
	  SELECT id, email, password_hash FROM users
	  WHERE email = ?;
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	userRow := stmt.QueryRowContext(ctx, email)
	var user models.User
	if err := userRow.Scan(&user.ID, &user.Email, &user.PasswordHashed); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passwordHash []byte) (int64, error) {
	const query = `INSERT INTO users(email, password_hash) VALUES(?, ?)`
	s.mu.Lock()
	defer s.mu.Unlock()
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, email, passwordHash)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const query = `SELECT is_admin FROM users WHERE id = ?;`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, err
	}
	res := stmt.QueryRowContext(ctx, userID)
	var isAdmin bool
	if err := res.Scan(&isAdmin); err != nil {
		return false, err
	}
	return isAdmin, nil
}
