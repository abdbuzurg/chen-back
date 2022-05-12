package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model `json:"-"`
	TableID    uint
	Status     bool `json:"status"`

	// One to Many
	OrderList []OrderList
}
