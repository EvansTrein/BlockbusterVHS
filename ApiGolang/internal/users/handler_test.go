package users_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/EvansTrein/BlockbusterVHS/internal/users"
	"github.com/EvansTrein/BlockbusterVHS/internal/users/mocks"
	"github.com/EvansTrein/BlockbusterVHS/logs"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	mockService := new(mocks.IUsersService)
	disLog := logs.NewDiscardLogger()

	handler := &users.HandlerUsers{
		Log:     disLog,
		Service: mockService,
	}

	t.Run("ServiceError_UserAlreadyExists", func(t *testing.T) {
		requestBody := users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Register", context.Background(), &requestBody).Return(nil, users.ErrUserAlreadyExsist).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to register user")
		assert.Contains(t, responseBody.Error, "user with this mail already exists")
	})

	t.Run("Success", func(t *testing.T) {
		requestBody := users.RegisterRequest{
			Name:     "Bob Varn",
			Email:    "bob@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		expectedResponse := &users.ReqisterResponce{
			ID: 1,
		}
		mockService.On("Register", context.Background(), &requestBody).Return(expectedResponse, nil).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, float64(1), responseBody["id"])
	})

	t.Run("InvalidBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`invalid-json`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to convert request body to json")
	})

	t.Run("ValidationError", func(t *testing.T) {
		requestBody := users.RegisterRequest{
			Name:     "",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "request body data failed validation")
	})

	t.Run("ServiceError_Timeout", func(t *testing.T) {
		requestBody := users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Register", context.Background(), &requestBody).Return(nil, context.DeadlineExceeded).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusGatewayTimeout, responseBody.Status)
		assert.Contains(t, responseBody.Message, "request processing exceeded the allowed time limit")
	})

	t.Run("ServiceError_InternalServerError", func(t *testing.T) {
		requestBody := users.RegisterRequest{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Register", context.Background(), &requestBody).Return(nil, errors.New("internal server error")).Once()

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Register()(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to register user")
		assert.Contains(t, responseBody.Error, "internal server error")
	})
}

func TestUpdateHandler(t *testing.T) {
	mockService := new(mocks.IUsersService)
	disLog := logs.NewDiscardLogger()

	handler := &users.HandlerUsers{
		Log:     disLog,
		Service: mockService,
	}

	t.Run("Success", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		expectedResponse := &users.UpdateResponce{
			ID:    1,
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "+300023143",
		}
		mockService.On("Update", context.Background(), &requestBody).Return(expectedResponse, nil).Once()

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, float64(1), responseBody["id"])
		assert.Equal(t, "John Doe", responseBody["name"])
		assert.Equal(t, "john.doe@example.com", responseBody["email"])
		assert.Equal(t, "+300023143", responseBody["phone"])
	})

	t.Run("InvalidBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBufferString(`invalid-json`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to convert request body to json")
	})

	t.Run("ValidationFailed", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       1,
			Name:     "",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "request body data failed validation")
	})

	t.Run("ServiceError_UserNotFound", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       999,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Update", context.Background(), &requestBody).Return(nil, users.ErrUserIdNotExist).Once()

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to update user")
		assert.Contains(t, responseBody.Error, "user with this ID does not exist")
	})

	t.Run("ServiceError_UserAlreadyExists", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Update", context.Background(), &requestBody).Return(nil, users.ErrUserAlreadyExsist).Once()

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to update user")
		assert.Contains(t, responseBody.Error, "user with this mail already exists")
	})

	t.Run("ServiceError_Timeout", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Update", context.Background(), &requestBody).Return(nil, context.DeadlineExceeded).Once()

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusGatewayTimeout, responseBody.Status)
		assert.Contains(t, responseBody.Message, "request processing exceeded the allowed time limit")
	})

	t.Run("ServiceError_InternalServerError", func(t *testing.T) {
		requestBody := users.UpdateRequest{
			ID:       1,
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Phone:    "+300023143",
			Password: "password123",
		}
		bodyBytes, _ := json.Marshal(requestBody)

		mockService.On("Update", context.Background(), &requestBody).Return(nil, errors.New("internal server error")).Once()

		req := httptest.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Update()(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to update user")
		assert.Contains(t, responseBody.Error, "internal server error")
	})
}

func TestGetUserHandler(t *testing.T) {
	mockService := new(mocks.IUsersService)
	disLog := logs.NewDiscardLogger()

	handler := &users.HandlerUsers{
		Log:     disLog,
		Service: mockService,
	}

	t.Run("Success", func(t *testing.T) {
		id := 1
		expectedResponse := users.GetUserResponce{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Phone: "+300023143",
		}
		mockService.On("User", context.Background(), uint(id)).Return(&expectedResponse, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(id), nil)
		w := httptest.NewRecorder()

		handler.GetUser()(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "John Doe", responseBody["name"])
		assert.Equal(t, "john.doe@example.com", responseBody["email"])
		assert.Equal(t, "+300023143", responseBody["phone"])
	})

	t.Run("InvalidID", func(t *testing.T) {
		idStr := "invalid"
		req := httptest.NewRequest(http.MethodGet, "/users/"+idStr, nil)
		w := httptest.NewRecorder()

		handler.GetUser()(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed to extract user id")
	})

	t.Run("ServiceError_UserNotFound", func(t *testing.T) {
		id := 999999
		mockService.On("User", context.Background(), uint(id)).Return(nil, users.ErrUserIdNotExist).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(id), nil)
		w := httptest.NewRecorder()

		handler.GetUser()(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed get user")
		assert.Contains(t, responseBody.Error, "user with this ID does not exist")
	})

	t.Run("ServiceError_Timeout", func(t *testing.T) {
		id := 1
		mockService.On("User", context.Background(), uint(id)).Return(nil, context.DeadlineExceeded).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(id), nil)
		w := httptest.NewRecorder()

		handler.GetUser()(w, req)

		assert.Equal(t, http.StatusGatewayTimeout, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusGatewayTimeout, responseBody.Status)
		assert.Contains(t, responseBody.Message, "request processing exceeded the allowed time limit")
	})

	t.Run("ServiceError_InternalServerError", func(t *testing.T) {
		id := 1
		mockService.On("User", context.Background(), uint(id)).Return(nil, errors.New("internal server error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/users/"+strconv.Itoa(id), nil)
		w := httptest.NewRecorder()

		handler.GetUser()(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var responseBody users.HandlerResponce
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, responseBody.Status)
		assert.Contains(t, responseBody.Message, "failed get user data")
		assert.Contains(t, responseBody.Error, "internal server error")
	})
}
