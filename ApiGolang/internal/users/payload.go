package users

type HandlerResponce struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
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
	ID       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"` // e164 formatted phone number
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateResponce struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetUserResponce struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
