package model

type Branch struct {
	OwnModel
	OrganizationID uint
	Name           string `json:"name"`
	IsActive       string `json:"is_active" gorm:"column:isActive"`

	//Many to Many
	Users []User `json:"-" gorm:"many2many:branches_users;"`
	// One To Many
	Halls []Hall `json:"-"`
}

type BranchCreateData struct {
	OrganizationID uint   `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	IsActive       string `json:"is_active" binding:"required"`
}

type BranchUpdateData struct {
	Name     string `json:"name" binding:"required"`
	IsActive string `json:"is_active" binding:"required"`
}
