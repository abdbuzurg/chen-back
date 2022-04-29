package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	BranchID   uint   `json:"-"`

	// One to Many
	Tables []Table `json:"-"`
}
