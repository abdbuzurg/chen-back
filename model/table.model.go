package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model `json:"-"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Number     uint    `json:"number"`
	HallID     uint
}
