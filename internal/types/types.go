package types

type Movie struct {
	Title string `json:"title" validate:"required"`
	Year  int    `json:"year" validate:"required,min=1888,max=2024"` // Movies start from 1888
}
