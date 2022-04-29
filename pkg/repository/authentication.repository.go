package repository

import (
	"chen/model"
	"chen/pkg/dto"

	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Create(registrationData dto.AuthenticationRegisterDTO) error
	UserFindByUsername(username string) (model.User, error)
}

type authenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return authenticationRepository{
		db: db,
	}
}

func (ar authenticationRepository) Create(registrationData dto.AuthenticationRegisterDTO) error {
	user := model.User{
		Username:  registrationData.Username,
		Password:  registrationData.Password,
		Firstname: registrationData.Firstname,
		Lastname:  registrationData.Lastname,
		RoleID:    registrationData.RoleID,
		IsActive:  true,
	}

	result := ar.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ar authenticationRepository) UserFindByUsername(username string) (model.User, error) {
	var user model.User
	result := ar.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
