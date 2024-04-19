package app

import (
	"context"
	"errors"

	"github.com/SashaMelva/tourism/internal/storage/memory"
	"github.com/SashaMelva/tourism/internal/storage/model"
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

func (a *App) GetAllUsers(ctx context.Context) ([]model.User, error) {
	event, err := a.storage.GetAllUseres()

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (a *App) RegisterUser(user model.User) (uint32, error) {
	if user.Password != user.RepeatPassword {
		return 0, errors.New("Repeate passwort")
	}

	createdUserId, err := a.storage.CreateUser(user)

	if err != nil {
		return 0, err
	}

	return createdUserId, err
}
