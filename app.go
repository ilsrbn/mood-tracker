package main

import (
	"context"
	"mood-tracker/src/core"
	"mood-tracker/src/rating"
	"mood-tracker/src/rating/repository"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (A *App) startup(ctx context.Context) {
	A.Ctx = ctx
	core.WithDB(repository.CreateRatingTable)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return core.Hello(name)
}

func (a *App) CreateRating(value int) (int64, error) {
	return rating.AddRating(value)
}

func (a *App) GetAllRatings() ([]repository.Rating, error) {
	return rating.GetAllRatings()
}
