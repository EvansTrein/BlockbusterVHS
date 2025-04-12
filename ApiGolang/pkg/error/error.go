package error

import "errors"

var (
	ErrUserAlreadyExsist = errors.New("user with this mail already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrModeSearch        = errors.New("incorrect user search mode")
)
