package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	IsActive   bool   `json:"is_active"`

	//One to Many
	Branches []Branch `json:"-"`

	//Many to Many
	Users []User `json:"-" gorm:"many2many:organizations_users;"`
}
