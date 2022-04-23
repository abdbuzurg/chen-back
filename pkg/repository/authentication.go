package repository

import (
	"chen/db"
	"chen/model"
)

func UserCreate(registrationData model.RegisterData) error {
	db := db.GetSQLiteConnection()
	user := model.User{
		Username:  registrationData.Username,
		Password:  registrationData.Password,
		FirstName: registrationData.FirstName,
		LastName:  registrationData.LastName,
		RoleID:    registrationData.RoleID,
		IsActive:  true,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UserFindByUsername(username string) (model.User, error) {
	db := db.GetSQLiteConnection()

	var user model.User
	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
