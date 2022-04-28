package model

type Hall struct {
	OwnModel
	Name     string `json:"name"`
	BranchID uint   `json:"-"`

	// One to Many
	Tables []Table `json:"-"`
}

type HallCreateData struct {
	Name     string `json:"name" binding:"required"`
	BranchID uint   `json:"branch_id" binding:"required"`
}

type HallUpdateData struct {
	Name string `json:"name" binding:"required"`
}
