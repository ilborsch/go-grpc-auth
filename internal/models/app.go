package models

import "context"

type App struct {
	ID     int64
	Name   string
	Secret string
}

type AppProvider interface {
	App(ctx context.Context, id int64) (App, error)
}
