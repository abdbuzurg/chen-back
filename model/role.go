package model

type Role struct {
	OwnModel
	Name string `json:"name"`

	// one to many with USER
	Users []User

	// MANY TO MANY with PERMISSION
	Permissions []Permission `gorm:"many2many:roles_permissions;"`
}
