package model

import "gorm.io/gorm"

type OrderList struct {
	gorm.Model `json:"-"`
	Count      uint
	Comment    string `gorm:"type:text"`
	OrderID    uint
	ItemID     uint
}
