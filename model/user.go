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

// For Registration
type RegisterData struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	RoleID    uint   `json:"role_id" binding:"required"`
}

// For Login
type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
