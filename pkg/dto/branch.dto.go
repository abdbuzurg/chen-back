package dto

type BranchCreate struct {
	OrganizationID uint   `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	IsActive       bool   `json:"is_active" binding:"required"`
}

type BranchUpdate struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
