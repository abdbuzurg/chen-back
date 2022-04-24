package model

type Hall struct {
	OwnModel
	Name     string `json:"namme"`
	BranchID uint

	// One to Many
	Tables []Table
}
