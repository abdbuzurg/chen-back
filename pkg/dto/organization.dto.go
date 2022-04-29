package dto

type OrganizationDTO struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
