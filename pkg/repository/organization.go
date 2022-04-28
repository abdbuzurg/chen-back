package repository

import (
	"chen/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type OrganizationRepo interface {
	FindAll() ([]model.Organization, error)
	FindById(id int) (model.Organization, error)
	Create(data model.OrganizationData) error
	Update(id int, data model.OrganizationData) error
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

	var org model.Organization
	result := or.db.First(&org, "id = ?", id)
	if result.Error != nil {
		return model.Organization{}, errors.New("org not found")
	}

	return org, nil
}
func (or organizationRepo) Create(data model.OrganizationData) error {

	newOrg := &model.Organization{
		IsActive: data.IsActive,
		Name:     data.Name,
	}

	result := or.db.Create(newOrg)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (or organizationRepo) Update(id int, data model.OrganizationData) error {

	var org model.Organization
	result := or.db.First(&org, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = or.db.Model(org).Updates(data).Update("updated_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (or organizationRepo) Delete(id int) error {
	return or.db.Delete(&model.Organization{}, id).Error
}
