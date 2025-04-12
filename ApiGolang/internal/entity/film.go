package entity

type Film struct {
	Title string `json:"title" validate:"required"`
	ID    int    `json:"id" validate:"required"`
}

func NewFilm() *Film {
	return &Film{}
}
