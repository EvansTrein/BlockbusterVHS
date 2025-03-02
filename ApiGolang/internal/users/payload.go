package users

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"` // e164 formatted phone number
	Password string `json:"password" validate:"required,min=8"`
}

type ReqisterResponce struct {
	ID uint `json:"user_id"`
}

type HandlerResponce struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
