package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model `json:"-"`
	Method     string `json:"method"`
	Path       string `json:"path"`
}
