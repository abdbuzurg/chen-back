package model

type Table struct {
	OwnModel
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Number uint    `json:"number"`
	HallID uint
}
