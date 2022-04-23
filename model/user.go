package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RoleID    uint
	IsActive  bool
}

// For Registration
type RegisterData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RoleID    uint   `json:"roleID"`
}

// For Login
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
