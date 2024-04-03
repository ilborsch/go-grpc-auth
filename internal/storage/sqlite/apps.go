package storage

import (
	"context"
	"go-grpc-auth/internal/models"
)

func (s *Storage) App(ctx context.Context, id int64) (models.App, error) {
	const query = `SELECT id, name, secret FROM apps WHERE id = ?;`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return models.App{}, err
	}
	res := stmt.QueryRowContext(ctx, id)
	var app models.App
	if err := res.Scan(&app.ID, &app.Name, &app.Secret); err != nil {
		return models.App{}, err
	}
	return app, nil
}
