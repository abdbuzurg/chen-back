package model

type Branch struct {
	OwnModel
	OrganizationID uint
	Name           string `json:"name"`
	IsActive       bool   `json:"is_active" gorm:"column:isActive"`

	//Many to Many
	Users []User `json:"-" gorm:"many2many:branches_users;"`
	// One To Many
	Halls []Hall `json:"-"`
}
