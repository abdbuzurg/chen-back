package dto

type BranchCreateDTO struct {
	OrganizationID uint   `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	IsActive       bool   `json:"is_active" binding:"required"`
}

type BranchUpdateDTO struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
