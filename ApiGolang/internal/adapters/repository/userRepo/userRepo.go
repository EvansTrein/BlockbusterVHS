package userRepo

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/EvansTrein/BlockbusterVHS/internal/dto"
	"github.com/EvansTrein/BlockbusterVHS/internal/entity"
	"github.com/EvansTrein/BlockbusterVHS/pkg/db/sqlite"
	myErr "github.com/EvansTrein/BlockbusterVHS/pkg/error"
)

const tableName = "users"

type UsersRepo struct {
	log *slog.Logger
	db  *sqlite.SqliteDB
}

type UsersRepoDeps struct {
	*slog.Logger
	*sqlite.SqliteDB
}

func NewUsersRepo(deps *UsersRepoDeps) *UsersRepo {
	return &UsersRepo{
		log: deps.Logger,
		db:  deps.SqliteDB,
	}
}

func (r *UsersRepo) Create(ctx context.Context, data *dto.UserCreateRequest) (int, error) {
	op := "Database: create user"
	log := r.log.With(slog.String("operation", op))
	log.Debug("Create func call", "data", data)

	query := `
		INSERT INTO clients (name, email, phone)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, data.Name, data.Email, data.Phone)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			log.Warn("failed to create a record in the database, mail already exists", "error", err)
			return 0, myErr.ErrUserAlreadyExsist
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
	return int(id), nil
}

func (r *UsersRepo) Find(ctx context.Context, param *dto.UserRequest) (*entity.User, error) {
	op := "Database: find user"
	log := r.log.With(slog.String("operation", op))
	log.Debug("Find func call", "param", param)

	var user entity.User

	switch param.Mode {
	case "email":
		query := fmt.Sprintf(`
		SELECT id, name, email, phone
		FROM %s 
		WHERE email = ?
		`, tableName)

		if err := r.db.QueryRowContext(ctx, query, param.Email).Scan(&user.ID, &user.Name, &user.Email); err != nil {
			if strings.Contains(err.Error(), "no rows") {
				log.Warn("user not found")
				return nil, myErr.ErrUserNotFound
			}
			log.Error("failed to execute the database user search request", "error", err)
			return nil, err
		}
	case "id":
		query := fmt.Sprintf(`
		SELECT id, name, email, phone
		FROM %s 
		WHERE id = ?
		`, tableName)

		if err := r.db.QueryRowContext(ctx, query, param.ID).Scan(&user.ID, &user.Name, &user.Email); err != nil {
			if strings.Contains(err.Error(), "no rows") {
				log.Warn("user not found")
				return nil, myErr.ErrUserNotFound
			}
			log.Error("failed to execute the database user search request", "error", err)
			return nil, err
		}
	default:
		log.Error("incorrect user search mode")
		return nil, myErr.ErrUserModeSearch
	}

	log.Info("user successfully found")
	return &user, nil
}

func (r *UsersRepo) Update(ctx context.Context, user *entity.User) error {
	// TODO:
	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, id int) error {
	// TODO:
	return nil
}
