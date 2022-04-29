package model

type Item struct {
	OwnModel
	Price float32 `json:"price"`
	Name  string  `json:"name"`

	// One to Many
	OrderList []OrderList
}
