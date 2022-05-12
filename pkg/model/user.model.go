package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" gorm:"unique; type:varchar(20)"`
	Password   string `gorm:"text"`
	Firstname  string `json:"firstname" gorm:"type:varchar(20)"`
	Lastname   string `json:"lastname" gorm:"type:varchar(20)"`
	RoleID     uint
	IsActive   bool
}
