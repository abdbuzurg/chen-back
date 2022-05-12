package model

import "gorm.io/gorm"

type Hall struct {
	gorm.Model `json:"-"`
	Name       string `json:"name" gorm:"type:varchar(100)"`
	BranchID   uint   `json:"-"`

	// One to Many
	Tables []Table `json:"-"`
}
