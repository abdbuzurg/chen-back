package repository

import (
	"chen/model"
	"chen/pkg/dto"

	"gorm.io/gorm"
)

type OrganizationRepo interface {
	FindAll() ([]model.Organization, error)
	FindById(id int) (model.Organization, error)
	Create(data dto.OrganizationDTO) error
	Update(id int, data dto.OrganizationDTO) error
	Delete(id int) error
}

type organizationRepo struct {
	db *gorm.DB
}

func NewOrganizationRepo(db *gorm.DB) OrganizationRepo {
	return organizationRepo{
		db: db,
	}
}

func (or organizationRepo) FindAll() ([]model.Organization, error) {
	orgs := []model.Organization{}
	err := or.db.Find(&orgs).Error
	return orgs, err
}

func (or organizationRepo) FindById(id int) (model.Organization, error) {
	org := model.Organization{}
	err := or.db.First(&org, "id = ?", id).Error
	return org, err
}

func (or organizationRepo) Create(data dto.OrganizationDTO) error {
	return or.db.Create(&model.Organization{
		Name:     data.Name,
		IsActive: data.IsActive,
	}).Error
}

func (or organizationRepo) Update(id int, data dto.OrganizationDTO) error {

	org := model.Organization{}
	err := or.db.First(&org, "id = ?", id).Error
	if err != nil {
		return err
	}

	return or.db.Model(org).Updates(data).Error
}

func (or organizationRepo) Delete(id int) error {
	return or.db.Delete(&model.Organization{}, id).Error
}
