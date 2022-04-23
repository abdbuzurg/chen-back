package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
