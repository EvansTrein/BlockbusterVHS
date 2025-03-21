package users

import (
	"context"
	"log/slog"
	"strings"

	"github.com/EvansTrein/BlockbusterVHS/internal/storages/postgres"
)

type UsersRepoPostgres struct {
	log  *slog.Logger
	repo *postgres.PostgresDB
}

type UsersRepoPostgresDeps struct {
	*slog.Logger
	*postgres.PostgresDB
}

func NewUsersRepoPostgres(deps *UsersRepoPostgresDeps) *UsersRepoPostgres {
	return &UsersRepoPostgres{
		log:  deps.Logger,
		repo: deps.PostgresDB,
	}
}

func (r *UsersRepoPostgres) Create(ctx context.Context, data *RegisterRequest) (uint, error) {
	op := "Database: create user"
	log := r.log.With(slog.String("operation", op))
	log.Debug("Create func call", "data", data)

	query := `
		INSERT INTO clients (username, email, phone, password_hash)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	var id uint
	err := r.repo.DB.QueryRow(ctx, query, data.Name, data.Email, data.Phone, data.Password).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			log.Warn("failed to create a record in the database, mail already exists", "error", err.Error())
			return 0, ErrUserAlreadyExsist
		}
		log.Error("failed to create a record in the database", "error", err.Error())
		return 0, err
	}

	log.Info("user successfully created")
	return id, nil
}

func (r *UsersRepoPostgres) ExistsByID(ctx context.Context, id uint) error {
	op := "Database: user existence check"
	log := r.log.With(slog.String("operation", op))
	log.Debug("ExistsByID func call", "id", id)

	query := `SELECT EXISTS (
		SELECT 1
		FROM clients
		WHERE id = $1
	);`

	var exists bool
	if err := r.repo.DB.QueryRow(ctx, query, id).Scan(&exists); err != nil {
		log.Error("failed to check if the user exists in the database", "error", err.Error())
		return err
	}

	if !exists {
		log.Warn("no user with this id")
		return ErrUserIdNotExist
	}

	log.Info("user with this ID was successfully found")
	return nil
}

func (r *UsersRepoPostgres) Update(ctx context.Context, data *UpdateRequest) error {
	op := "Database: user update"
	log := r.log.With(slog.String("operation", op))
	log.Debug("Update func call", "data", data)

	query := `
		UPDATE clients
		SET username = $1, email = $2, phone = $3, password_hash = $4
		WHERE id = $5;
	`

	if _, err := r.repo.DB.Exec(ctx, query, data.Name, data.Email, data.Phone, data.Password, data.ID); err != nil {
		log.Error("failed to update database record")
		return err
	}

	log.Info("user successfully updated")
	return nil
}
