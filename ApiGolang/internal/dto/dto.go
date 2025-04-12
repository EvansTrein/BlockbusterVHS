package dto

type UserCreateRequest struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,e164"` // e164 formatted phone number
}

type UserCreateResponce struct {
	ID int `json:"id"`
}

type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,e164"`
	ID    int    `json:"id" validate:"required"`
}

type UserRequest struct {
	Mode  string
	Email string
	ID    int
}
