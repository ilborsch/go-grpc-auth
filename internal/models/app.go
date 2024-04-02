package models

import "context"

type App struct {
	ID     int
	Name   string
	Secret string
}

type AppProvider interface {
	App(ctx context.Context, id int) (App, error)
}
