package useCase

import (
	"context"
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/internal/dto"
	"github.com/EvansTrein/BlockbusterVHS/internal/entity"
)

//go:generate mockgen -source=$GOFILE -destination=../mocks/mock_userRepo.go -package=mocksRepository
type IUserRepository interface {
	Create(ctx context.Context, user *dto.UserCreateRequest) (int, error)
	Find(ctx context.Context, param *dto.UserRequest) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
}

type UserUseCase struct {
	log  *slog.Logger
	repo IUserRepository
}

type UserUseCaseDeps struct {
	*slog.Logger
	IUserRepository
}

func NewUserUseCase(deps *UserUseCaseDeps) *UserUseCase {
	return &UserUseCase{repo: deps.IUserRepository, log: deps.Logger}
}

func (uc *UserUseCase) Create(ctx context.Context, data *dto.UserCreateRequest) (*dto.UserCreateResponce, error) {
	op := "usecase user: user creation"
	log := uc.log.With(slog.String("operation", op))
	log.Debug("Create func call", "data", data)

	// TODO:

	log.Info("user successfully created")
	return nil, nil
}

func (uc *UserUseCase) User(ctx context.Context, id int) (*entity.User, error) {
	op := "usecase user: user reception"
	log := uc.log.With(slog.String("operation", op))
	log.Debug("User func call", "id", id)

	user, err := uc.repo.Find(ctx, &dto.UserRequest{Mode: "id", ID: id})
	if err != nil {
		return nil, err
	}

	log.Info("user successfully received")
	return user, nil
}

func (uc *UserUseCase) Update(ctx context.Context, user *entity.User) error {
	op := "usecase user: user data update"
	log := uc.log.With(slog.String("operation", op))
	log.Debug("Update func call", "user", user)

	if err := uc.repo.Update(ctx, user); err != nil {
		return err
	}

	log.Info("user successfully updated")
	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, id int) error {
	op := "usecase user: user deletion"
	log := uc.log.With(slog.String("operation", op))
	log.Debug("Delete func call", "id", id)

	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}

	log.Info("user successfully deleted")
	return nil
}
