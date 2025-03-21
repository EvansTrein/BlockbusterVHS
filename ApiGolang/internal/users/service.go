package users

import (
	"context"
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/pkg/utils"
)

type IUsersRepo interface {
	Create(ctx context.Context, data *RegisterRequest) (uint, error)
	ExistsByID(ctx context.Context, id uint) error
	Update(ctx context.Context, data *UpdateRequest) error
}

type UsersService struct {
	log *slog.Logger
	db  IUsersRepo
}

type UsersServiceDeps struct {
	*slog.Logger
	IUsersRepo
}

func NewUsersService(deps *UsersServiceDeps) *UsersService {
	return &UsersService{
		log: deps.Logger,
		db:  deps.IUsersRepo,
	}
}

func (s *UsersService) Register(ctx context.Context, data *RegisterRequest) (*ReqisterResponce, error) {
	op := "service Users: user registration started"
	log := s.log.With(slog.String("operation", op))
	log.Debug("Register func call", "requets data", data)

	hashPassword, err := utils.Hashing(data.Password)
	if err != nil {
		log.Error("failed to hash the password", "error", err)
		return nil, err
	}

	data.Password = hashPassword

	result, err := s.db.Create(ctx, data)
	if err != nil {
		log.Error("failed to register user", "error", err)
		return nil, err
	}

	log.Info("user successfully registered")
	return &ReqisterResponce{ID: result}, nil
}

func (s *UsersService) Update(ctx context.Context, data *UpdateRequest) (*UpdateResponce, error) {
	op := "service Users: user update started"
	log := s.log.With(slog.String("operation", op))
	log.Debug("Update func call", "requets data", data)

	if err := s.db.ExistsByID(ctx, data.ID); err != nil {
		log.Warn("failed to find user")
		return nil, err
	}

	hashPassword, err := utils.Hashing(data.Password)
	if err != nil {
		log.Error("failed to hash the password", "error", err)
		return nil, err
	}

	data.Password = hashPassword

	if err := s.db.Update(ctx, data); err != nil {
		log.Error("failed to updated user", "error", err)
		return nil, err
	}

	log.Info("user successfully updated")
	return &UpdateResponce{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Phone: data.Phone,
	}, nil
}
