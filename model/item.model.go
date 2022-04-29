package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model `json:"-"`
	Price      float32 `json:"price"`
	Name       string  `json:"name"`

	// One to Many
	OrderList []OrderList
}
