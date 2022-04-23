package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name"`

	// one to many with USER
	Users []User

	// MANY TO MANY with PERMISSION
	Permissions []Permission `gorm:"many2many:roles_permissions;"`
}
