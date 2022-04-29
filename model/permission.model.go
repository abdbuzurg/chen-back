package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model  `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
