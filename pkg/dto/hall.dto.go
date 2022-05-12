package dto

type HallCreate struct {
	Name     string `json:"name" binding:"required"`
	BranchID uint   `json:"branch_id" binding:"required"`
}

type HallUpdate struct {
	Name string `json:"name" binding:"required"`
}
