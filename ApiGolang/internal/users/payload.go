package users

type HandlerResponce struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"` // e164 formatted phone number
	Password string `json:"password" validate:"required,min=8"`
}

type ReqisterResponce struct {
	ID uint `json:"id"`
}

type UpdateRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=8"`
	ID       uint   `json:"id" validate:"required"`
}

type UpdateResponce struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	ID    uint   `json:"id"`
}

type GetUserResponce struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
