package sqlite

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	db  *sql.DB
	log *slog.Logger
}

func New(storagePath string, log *slog.Logger) (*SqliteDB, error) {
	log.Debug("database: connection to Postgres started")

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Info("database: connect to Postgres successfully")
	return &SqliteDB{db: db, log: log}, nil
}

func (s *SqliteDB) Close() error {
	s.log.Debug("database: stop started")

	if s.db == nil {
		return fmt.Errorf("database connection is already closed")
	}

	s.db.Close()

	s.db = nil

	s.log.Info("database: stop successful")
	return nil
}
