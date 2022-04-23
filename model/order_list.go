package model

import "gorm.io/gorm"

type OrderList struct {
	gorm.Model
	Count   uint
	OrderID uint
	ItemID  uint
}
