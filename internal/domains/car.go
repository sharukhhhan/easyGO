package domains

type Car struct {
	Id              int `json:"id"`
	CarMarkId       int `json:"carMarkId"`
	ManufactureYear int `json:"manufactureYear"`
}

type CarMark struct {
	Id       int    `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Capacity int    `json:"capacity"`
}
