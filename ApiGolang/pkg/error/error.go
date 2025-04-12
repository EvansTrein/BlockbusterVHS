package error

import "errors"

var (
	ErrUserAlreadyExsist = errors.New("user with this mail already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserModeSearch    = errors.New("incorrect user search mode")
)
