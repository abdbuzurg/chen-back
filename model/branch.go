package model

type Branch struct {
	OwnModel
	OrganizationID uint
	Name           string `json:"name"`
	IsActive       string `json:"isActive" gorm:"column:isActive"`

	//Many to Many
	Users []User `gorm:"many2many:branches_users;"`

	// One To Many
	Halls []Hall
}
