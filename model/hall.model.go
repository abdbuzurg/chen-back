package model

type Hall struct {
	OwnModel
	Name     string `json:"name"`
	BranchID uint   `json:"-"`

	// One to Many
	Tables []Table `json:"-"`
}
