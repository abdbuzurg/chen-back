package model

type Permission struct {
	OwnModel
	Title       string `json:"title"`
	Description string `json:"description"`
}
