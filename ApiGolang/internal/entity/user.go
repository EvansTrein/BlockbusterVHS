package entity

type User struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required,e164"`
	ID    int    `json:"id" validate:"required"`
}

func NewUser() *User {
	return &User{}
}
