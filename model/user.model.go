package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username"`
	Password   string
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	RoleID     uint
	IsActive   bool
}
