package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB  *sql.DB
	log *slog.Logger
}

func New(storagePath string, log *slog.Logger) (*SqliteDB, error) {
	log.Debug("database: connection to SQLite started")

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info("database: connect to SQLite successfully")
	return &SqliteDB{DB: db, log: log}, nil
}

func (s *SqliteDB) Close() error {
	s.log.Debug("database: stop started")

	if s.DB == nil {
		return errors.New("database connection is already closed")
	}

	s.DB.Close()

	s.DB = nil

	s.log.Info("database: stop successful")
	return nil
}
