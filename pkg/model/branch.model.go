package model

import "gorm.io/gorm"

type Branch struct {
	gorm.Model     `json:"-"`
	OrganizationID uint
	Name           string `json:"name" gorm:"varchar(100)"`
	IsActive       bool   `json:"is_active" gorm:"column:isActive"`

	//Many to Many
	Users []User `json:"-" gorm:"many2many:branches_users;"`
	// One To Many
	Halls []Hall `json:"-"`
}
