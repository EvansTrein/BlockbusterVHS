package httpAdapter

import (
	"context"
	"net/http"

	"github.com/EvansTrein/BlockbusterVHS/internal/dto"
	"github.com/EvansTrein/BlockbusterVHS/internal/entity"
)

type IUserUseCase interface {
	Create(ctx context.Context, data *dto.UserCreateRequest) (*dto.UserCreateResponce, error)
	User(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
}

type HandlerUser struct {
	baseH *BaseHandler
	uc    IUserUseCase
}

type HandlerUserDeps struct {
	*BaseHandler
	IUserUseCase
}

func NewHandlerUser(deps *HandlerUserDeps) *HandlerUser {
	return &HandlerUser{
		baseH: deps.BaseHandler,
		uc:    deps.IUserUseCase,
	}
}

func (h *HandlerUser) UserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *HandlerUser) User() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// userID, ok := r.Context().Value(UserIDKey).(int)
		// if !ok {
		// 	h.baseH.HandleError(w, myErr.ErrNoUserID)
		// 	return
		// }

		// resp, err := h.userUC.User(r.Context(), userID)
		// if err != nil {
		// 	h.baseH.HandleError(w, err)
		// 	return
		// }

		// h.baseH.SendJsonResp(w, 200, resp)
	}
}

func (h *HandlerUser) UserUpd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// userID, ok := r.Context().Value(UserIDKey).(int)
		// if !ok {
		// 	h.baseH.HandleError(w, myErr.ErrNoUserID)
		// 	return
		// }

		// data, err := h.baseH.Handle(w, r, func(body io.Reader) (any, error) {
		// 	return utils.DecodeBody[entity.User](r.Body)
		// })
		// if err != nil {
		// 	return
		// }

		// updData, ok := data.(*entity.User)
		// if !ok {
		// 	h.baseH.HandleError(w, myErr.ErrTypeConversion)
		// 	return
		// }

		// if userID != updData.ID {
		// 	h.baseH.HandleError(w, myErr.ErrIdsNotMatch)
		// 	return
		// }

		// if err := h.userUC.UserUpd(r.Context(), updData); err != nil {
		// 	h.baseH.HandleError(w, myErr.ErrIdsNotMatch)
		// 	return
		// }

		// w.WriteHeader(204)
	}
}

func (h *HandlerUser) UserDel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// sessionID, ok := r.Context().Value(SessionIDKey).(string)
		// if !ok {
		// 	h.baseH.HandleError(w, myErr.ErrNoSessionID)
		// 	return
		// }

		// userID, ok := r.Context().Value(UserIDKey).(int)
		// if !ok {
		// 	h.baseH.HandleError(w, myErr.ErrNoUserID)
		// 	return
		// }

		// if err := h.userUC.UserDel(r.Context(), userID); err != nil {
		// 	h.baseH.HandleError(w, err)
		// 	return
		// }

		// if err := h.authUC.SessionDel(r.Context(), sessionID); err != nil {
		// 	h.baseH.HandleError(w, err)
		// 	return
		// }
		// cookieManager := &cookie.CookieManager{}
		// cookieManager.DeleteCookie(w, string(SessionIDKey))
		// h.baseH.Log.Info("user session successfully closed")

		// w.WriteHeader(204)
	}
}
