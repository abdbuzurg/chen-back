package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model `json:"-"`
	Price      float32 `json:"price"`
	Name       string  `json:"name" gorm:"unique; type:varchar(100)"`

	// One to Many
	OrderList []OrderList
}
