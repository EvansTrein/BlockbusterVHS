package filmRepo

import (
	"context"
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/pkg/db/sqlite"
)

// const tableName = "films"

type FilmsRepo struct {
	log *slog.Logger
	db  *sqlite.SqliteDB
}

type FilmsRepoDeps struct {
	*slog.Logger
	*sqlite.SqliteDB
}

func NewFilmsRepo(deps *FilmsRepoDeps) *FilmsRepo {
	return &FilmsRepo{
		log: deps.Logger,
		db:  deps.SqliteDB,
	}
}

func (r *FilmsRepo) Create(ctx context.Context) error {
	// TODO:
	return nil
}

func (r *FilmsRepo) Find(ctx context.Context) error {
	// TODO:
	return nil
}

func (r *FilmsRepo) Update(ctx context.Context) error {
	// TODO:
	return nil
}

func (r *FilmsRepo) Delete(ctx context.Context) error {
	// TODO:
	return nil
}
