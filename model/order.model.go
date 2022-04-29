package model

type Order struct {
	OwnModel
	TableID uint
	Status  bool `json:"status"`

	// One to Many
	OrderList []OrderList
}
