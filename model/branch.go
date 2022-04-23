package model

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	OrganizationID uint
	Name           string `json:"name"`
	IsActive       string `json:"isActive"`

	//Many to Many
	Users []User `gorm:"many2many:branches_users;"`

	// One To Many
	Halls []Hall
}
