package users

import "errors"

var (
	ErrUserAlreadyExsist = errors.New("user with this mail already exists")
)
