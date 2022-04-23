package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`

	//One to Many
	Branches []Branch

	//Many to Many
	Users []User `gorm:"many2many:organizations_users;"`
}

type OrganizationData struct {
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
}
