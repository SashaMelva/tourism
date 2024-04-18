package app

import (
	"context"

	"github.com/SashaMelva/tourism/internal/storage/memory"
	"go.uber.org/zap"
)

type App struct {
	storage *memory.Storage
	Logger  *zap.SugaredLogger
}

// type Storage interface {
// }

func New(logger *zap.SugaredLogger, storage *memory.Storage) *App {
	return &App{
		storage: storage,
		Logger:  logger,
	}
}

func (a *App) GetAllUsers(ctx context.Context) ([]memory.User, error) {
	event, err := a.storage.GetAllUseres()

	if err != nil {
		return nil, err
	}

	return event, nil
}
