package users_test

import (
	"context"
	"errors"
	"testing"

	"github.com/EvansTrein/BlockbusterVHS/internal/users"
	"github.com/EvansTrein/BlockbusterVHS/internal/users/mocks"
	"github.com/EvansTrein/BlockbusterVHS/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockRepo := new(mocks.IUsersRepo)
	disLog := logs.NewDiscardLogger()

	service := &users.UsersService{
		Log: disLog,
		Db:  mockRepo,
	}

	t.Run("Success", func(t *testing.T) {
		requestBody := &users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("Create", mock.Anything, requestBody).Return(uint(1), nil).Once()

		response, err := service.Register(context.Background(), requestBody)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)

		mockRepo.AssertExpectations(t)
	})

	t.Run("HashingError", func(t *testing.T) {
		requestBody := &users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}

		mockRepo.On("Create", mock.Anything, requestBody).Return(uint(0), errors.New("failed to hash the password")).Once()

		_, err := service.Register(context.Background(), requestBody)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to hash the password")
	})

	t.Run("DatabaseError", func(t *testing.T) {
		requestBody := &users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("Create", mock.Anything, requestBody).Return(uint(0), users.ErrUserAlreadyExsist).Once()

		_, err := service.Register(context.Background(), requestBody)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user with this mail already exists")

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockRepo := new(mocks.IUsersRepo)
	disLog := logs.NewDiscardLogger()

	service := &users.UsersService{
		Log: disLog,
		Db:  mockRepo,
	}

	t.Run("Success", func(t *testing.T) {
		requestBody := &users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("ExistsByID", mock.Anything, uint(1)).Return(nil).Once()
		mockRepo.On("Update", mock.Anything, requestBody).Return(nil).Once()

		response, err := service.Update(context.Background(), requestBody)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "John Doe", response.Name)

		mockRepo.AssertExpectations(t)
	})

	t.Run("HashingError", func(t *testing.T) {
		requestBody := &users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("ExistsByID", mock.Anything, uint(1)).Return(nil).Once()
		mockRepo.On("Update", mock.Anything, requestBody).Return(errors.New("failed to hash the password")).Once()

		_, err := service.Update(context.Background(), requestBody)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to hash the password")
	})

	t.Run("UserNotFound", func(t *testing.T) {
		requestBody := &users.UpdateRequest{
			ID:       999,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("ExistsByID", mock.Anything, uint(999)).Return(users.ErrUserIdNotExist).Once()

		_, err := service.Update(context.Background(), requestBody)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user with this ID does not exist")

		mockRepo.AssertExpectations(t)
	})

	t.Run("DatabaseError", func(t *testing.T) {
		requestBody := &users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		mockRepo.On("ExistsByID", mock.Anything, uint(1)).Return(nil).Once()
		mockRepo.On("Update", mock.Anything, requestBody).Return(errors.New("failed to updated user")).Once()

		_, err := service.Update(context.Background(), requestBody)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to updated user")

		mockRepo.AssertExpectations(t)
	})
}

func TestUser(t *testing.T) {
	mockRepo := new(mocks.IUsersRepo)
	disLog := logs.NewDiscardLogger()

	service := &users.UsersService{
		Log: disLog,
		Db:  mockRepo,
	}

	t.Run("Success", func(t *testing.T) {
		id := uint(1)
		expectedResponse := &users.GetUserResponce{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "+300023143",
		}
		mockRepo.On("ExistsByID", mock.Anything, id).Return(nil).Once()
		mockRepo.On("GetUserData", mock.Anything, id, mock.Anything).Run(func(args mock.Arguments) {
			resp := args.Get(2).(*users.GetUserResponce)
			*resp = *expectedResponse
		}).Return(nil).Once()

		response, err := service.User(context.Background(), id)

		assert.NoError(t, err)
		assert.Equal(t, expectedResponse.Name, response.Name)
		assert.Equal(t, expectedResponse.Email, response.Email)
		assert.Equal(t, expectedResponse.Phone, response.Phone)

		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		id := uint(999)
		mockRepo.On("ExistsByID", mock.Anything, id).Return(users.ErrUserIdNotExist).Once()

		_, err := service.User(context.Background(), id)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user with this ID does not exist")

		mockRepo.AssertExpectations(t)
	})

	t.Run("DatabaseError", func(t *testing.T) {
		id := uint(1)
		mockRepo.On("ExistsByID", mock.Anything, id).Return(nil).Once()
		mockRepo.On("GetUserData", mock.Anything, id, mock.Anything).Return(errors.New("failed to retrieve user data")).Once()

		_, err := service.User(context.Background(), id)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve user data")

		mockRepo.AssertExpectations(t)
	})
}
