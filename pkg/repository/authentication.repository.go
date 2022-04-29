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
	return ar.db.Create(&model.User{
		Username:  registrationData.Username,
		Password:  registrationData.Password,
		Firstname: registrationData.Firstname,
		Lastname:  registrationData.Lastname,
		RoleID:    registrationData.RoleID,
		IsActive:  true,
	}).Error
}

func (ar authenticationRepository) UserFindByUsername(username string) (model.User, error) {
	user := model.User{}
	err := ar.db.First(&user, "username = ?", username).Error
	return user, err
}
