package model

type User struct {
	OwnModel
	Username  string `json:"username"`
	Password  string
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	RoleID    uint
	IsActive  bool
}
