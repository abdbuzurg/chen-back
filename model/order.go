package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TableID uint
	Status  bool `json:"status"`

	// One to Many
	OrderList []OrderList
}
