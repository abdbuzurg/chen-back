package model

type Organization struct {
	OwnModel
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`

	//One to Many
	Branches []Branch `json:"-"`

	//Many to Many
	Users []User `json:"-" gorm:"many2many:organizations_users;"`
}

type OrganizationData struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
