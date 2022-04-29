package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model  `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`

	// one to many with USER
	Users []User `json:"-"`

	// MANY TO MANY with PERMISSION
	Permissions []Permission `json:"-" gorm:"many2many:roles_permissions;"`
}
