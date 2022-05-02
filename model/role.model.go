package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model     `json:"-"`
	Title          string `json:"title" gorm:"type:varchar(40)"`
	Description    string `json:"description"`
	OrganizationID uint   `json:"-"`

	// one to many with USER
	Users []User `json:"-"`

	// MANY TO MANY with PERMISSION
	Permissions []Permission `json:"-" gorm:"many2many:roles_permissions;"`
}
