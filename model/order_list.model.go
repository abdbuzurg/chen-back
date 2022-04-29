package model

import "gorm.io/gorm"

type OrderList struct {
	gorm.Model `json:"-"`
	Count      uint
	OrderID    uint
	ItemID     uint
}
