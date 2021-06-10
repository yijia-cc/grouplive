package entity

type Type struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Category *Category `json:"category"`
}
