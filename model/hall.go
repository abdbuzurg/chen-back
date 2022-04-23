package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model
	Name     string `json:"namme"`
	BranchID uint

	// One to Many
	Tables []Table
}
