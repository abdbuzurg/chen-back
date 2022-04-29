package dto

type HallCreateDTO struct {
	Name     string `json:"name" binding:"required"`
	BranchID uint   `json:"branch_id" binding:"required"`
}

type HallUpdateDTO struct {
	Name string `json:"name" binding:"required"`
}
