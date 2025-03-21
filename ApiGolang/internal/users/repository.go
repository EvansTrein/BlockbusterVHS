package users

import (
	"context"
	"log/slog"
	"strings"

	"github.com/EvansTrein/BlockbusterVHS/internal/storages/sqlite"
)

type UsersRepo struct {
	log  *slog.Logger
	repo *sqlite.SqliteDB
}

type UsersRepoDeps struct {
	*slog.Logger
	*sqlite.SqliteDB
}

func NewUsersRepo(deps *UsersRepoDeps) *UsersRepo {
	return &UsersRepo{
		log:  deps.Logger,
		repo: deps.SqliteDB,
	}
}

func (r *UsersRepo) Create(ctx context.Context, data *RegisterRequest) (uint, error) {
	op := "Database: create user"
	log := r.log.With(slog.String("operation", op))
	log.Debug("Create func call", "data", data)

	query := `
		INSERT INTO clients (username, email, phone, password_hash)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.repo.DB.ExecContext(ctx, query, data.Name, data.Email, data.Phone, data.Password)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			log.Warn("failed to create a record in the database, mail already exists", "error", err)
			return 0, ErrUserAlreadyExsist
		}
		log.Error("failed to create a record in the database", "error", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error("failed to get the id of the created user", "error", err)
		return 0, err
	}

	log.Info("user successfully created")
	return uint(id), nil
}
