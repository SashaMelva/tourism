package sqlstorage

import (
	"context"
	"database/sql"

	"github.com/SashaMelva/tourism/internal/config"
	_ "github.com/jackc/pgx/stdlib"
	"go.uber.org/zap"
)

type StorageConnection struct {
	StorageDb *sql.DB
}

func New(confDB *config.ConfigDB, log *zap.SugaredLogger) *StorageConnection {
	dsn := "postgres://" + confDB.User + ":" + confDB.Password + "@" + confDB.Host + ":" + confDB.Port + "/" + confDB.NameDB
	storage, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Cannot open pgx driver: %w", err)
	}

	return &StorageConnection{storage}
}

func (s *StorageConnection) Connect(ctx context.Context) error {
	err := s.StorageDb.PingContext(ctx)
	return err
}

func (s *StorageConnection) Close(ctx context.Context) error {
	s.StorageDb.Close()
	return nil
}
