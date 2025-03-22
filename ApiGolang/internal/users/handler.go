package users

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/EvansTrein/BlockbusterVHS/pkg/utils"
	"github.com/EvansTrein/BlockbusterVHS/pkg/validate"
)

//go:generate mockery --name IUsersService --output ./mocks --outpkg mocks
type IUsersService interface {
	Register(ctx context.Context, data *RegisterRequest) (*ReqisterResponce, error)
	Update(ctx context.Context, data *UpdateRequest) (*UpdateResponce, error)
	User(ctx context.Context, id uint) (*GetUserResponce, error)
}

type HandlerUsers struct {
	Log     *slog.Logger
	Service IUsersService
}

type HandlerUsersDeps struct {
	*slog.Logger
	IUsersService
}

func NewHandler(router *http.ServeMux, deps *HandlerUsersDeps) {
	handler := &HandlerUsers{
		Log:     deps.Logger,
		Service: deps.IUsersService,
	}

	router.HandleFunc("POST /users", handler.Register())
	router.HandleFunc("PUT /users", handler.Update())
	router.HandleFunc("GET /users/{id}", handler.GetUser())
}

func (u *HandlerUsers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "Handler Register: call"
		log := u.Log.With(
			slog.String("operation", op),
			slog.String("apiPath", r.URL.Path),
			slog.String("HTTP Method", r.Method),
		)
		log.Debug("request received")

		data, err := utils.DecodeBody[RegisterRequest](r.Body)
		if err != nil {
			log.Error("failed to convert request body to json", "error", err)
			utils.SendJsonResp(w, 400, &HandlerResponce{
				Status:  http.StatusBadRequest,
				Message: "failed to convert request body to json",
				Error:   err.Error(),
			})
			return
		}

		if err := validate.IsValid(data); err != nil {
			log.Error("request body data failed validation", "error", err)
			utils.SendJsonResp(w, 400, &HandlerResponce{
				Status:  http.StatusBadRequest,
				Message: "request body data failed validation",
				Error:   err.Error(),
			})
			return
		}

		log.Debug("data successfully validated", "data", data)

		resp, err := u.Service.Register(r.Context(), data)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserAlreadyExsist):
				log.Warn("failed to register user", "error", err)
				utils.SendJsonResp(w, 400, &HandlerResponce{
					Status:  http.StatusBadRequest,
					Message: "failed to register user",
					Error:   err.Error(),
				})
				return
			case errors.Is(err, context.DeadlineExceeded):
				log.Error("request processing exceeded the allowed time limit", "error", err)
				utils.SendJsonResp(w, 504, &HandlerResponce{
					Status:  http.StatusGatewayTimeout,
					Message: "request processing exceeded the allowed time limit",
					Error:   err.Error(),
				})
				return
			default:
				log.Error("failed to register user", "error", err)
				utils.SendJsonResp(w, 500, &HandlerResponce{
					Status:  http.StatusInternalServerError,
					Message: "failed to register user",
					Error:   err.Error(),
				})
				return
			}
		}

		utils.SendJsonResp(w, 201, resp)
		log.Info("response successfully sent")
	}
}

func (u *HandlerUsers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "Handler Update: call"
		log := u.Log.With(
			slog.String("operation", op),
			slog.String("apiPath", r.URL.Path),
			slog.String("HTTP Method", r.Method),
		)
		log.Debug("request received")

		data, err := utils.DecodeBody[UpdateRequest](r.Body)
		if err != nil {
			log.Error("failed to convert request body to json", "error", err)
			utils.SendJsonResp(w, 400, &HandlerResponce{
				Status:  http.StatusBadRequest,
				Message: "failed to convert request body to json",
				Error:   err.Error(),
			})
			return
		}

		if err := validate.IsValid(data); err != nil {
			log.Error("request body data failed validation", "error", err)
			utils.SendJsonResp(w, 400, &HandlerResponce{
				Status:  http.StatusBadRequest,
				Message: "request body data failed validation",
				Error:   err.Error(),
			})
			return
		}

		log.Debug("data successfully validated", "data", data)

		resp, err := u.Service.Update(r.Context(), data)
		if err != nil {
			switch {
			case errors.Is(err, ErrUserIdNotExist):
				log.Warn("failed to update user", "error", err)
				utils.SendJsonResp(w, 404, &HandlerResponce{
					Status:  http.StatusNotFound,
					Message: "failed to update user",
					Error:   err.Error(),
				})
				return
			case errors.Is(err, ErrUserAlreadyExsist):
				log.Warn("failed to register user", "error", err)
				utils.SendJsonResp(w, 400, &HandlerResponce{
					Status:  http.StatusBadRequest,
					Message: "failed to update user",
					Error:   err.Error(),
				})
				return
			case errors.Is(err, context.DeadlineExceeded):
				log.Error("request processing exceeded the allowed time limit", "error", err)
				utils.SendJsonResp(w, 504, &HandlerResponce{
					Status:  http.StatusGatewayTimeout,
					Message: "request processing exceeded the allowed time limit",
					Error:   err.Error(),
				})
				return
			default:
				log.Error("failed to update user", "error", err)
				utils.SendJsonResp(w, 500, &HandlerResponce{
					Status:  http.StatusInternalServerError,
					Message: "failed to update user",
					Error:   err.Error(),
				})
				return
			}
		}

		utils.SendJsonResp(w, 200, resp)
		log.Info("response successfully sent")
	}
}

func (u *HandlerUsers) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "Handler GetUser: call"
		log := u.Log.With(
			slog.String("operation", op),
			slog.String("apiPath", r.URL.Path),
			slog.String("HTTP Method", r.Method),
		)
		log.Debug("request received")

		id, err := utils.ExtractID(r.URL.Path, 2)
		if err != nil {
			log.Warn("failed to extract user id", "error", err)
			utils.SendJsonResp(w, 400, &HandlerResponce{
				Status:  http.StatusBadRequest,
				Message: "failed to extract user id",
				Error:   err.Error(),
			})
			return
		}

		log.Debug("data successfully validated", "id", id)

		resp, err := u.Service.User(r.Context(), uint(id))
		if err != nil {
			switch {
			case errors.Is(err, ErrUserIdNotExist):
				log.Warn("failed get user", "error", err)
				utils.SendJsonResp(w, 404, &HandlerResponce{
					Status:  http.StatusNotFound,
					Message: "failed get user",
					Error:   err.Error(),
				})
				return
			case errors.Is(err, context.DeadlineExceeded):
				log.Error("request processing exceeded the allowed time limit", "error", err)
				utils.SendJsonResp(w, 504, &HandlerResponce{
					Status:  http.StatusGatewayTimeout,
					Message: "request processing exceeded the allowed time limit",
					Error:   err.Error(),
				})
				return
			default:
				log.Error("failed get user data", "error", err)
				utils.SendJsonResp(w, 500, &HandlerResponce{
					Status:  http.StatusInternalServerError,
					Message: "failed get user data",
					Error:   err.Error(),
				})
				return
			}
		}

		utils.SendJsonResp(w, 200, resp)
		log.Info("response successfully sent")
	}
}
