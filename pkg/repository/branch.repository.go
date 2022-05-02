package repository

import (
	"chen/model"
	"chen/pkg/dto"

	"gorm.io/gorm"
)

type BranchRepository interface {
	FindAll() ([]model.Branch, error)
	FindById(id int) (model.Branch, error)
	Create(data dto.BranchCreateDTO) error
	Update(id int, data dto.BranchUpdateDTO) error
	Delete(id int) error
}

type branchRepositry struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return branchRepositry{
		db: db,
	}
}

func (br branchRepositry) FindAll() ([]model.Branch, error) {
	orgs := []model.Branch{}
	err := br.db.Find(&orgs).Error
	return orgs, err
}

func (br branchRepositry) FindById(id int) (model.Branch, error) {
	branch := model.Branch{}
	err := br.db.First(&branch, "id = ?", id).Error

	return branch, err
}

func (br branchRepositry) Create(data dto.BranchCreateDTO) error {
	return br.db.Create(&model.Branch{
		OrganizationID: data.OrganizationID,
		Name:           data.Name,
		IsActive:       data.IsActive,
	}).Error
}

func (br branchRepositry) Update(id int, data dto.BranchUpdateDTO) error {
	branch := model.Branch{}
	err := br.db.First(branch, "id = ?", id).Error
	if err != nil {
		return err
	}

	return br.db.Model(branch).Updates(data).Error
}

func (br branchRepositry) Delete(id int) error {
	return br.db.Delete(&model.Branch{}, "id = ?", id).Error
}
