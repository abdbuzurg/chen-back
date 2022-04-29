package repository

import (
	"chen/model"
	"chen/pkg/dto"
	"time"

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
	result := br.db.First(branch, "id = ?", id)
	if result.Error != nil {
		return model.Branch{}, result.Error
	}

	return branch, nil
}

func (br branchRepositry) Create(data dto.BranchCreateDTO) error {

	newBranch := &model.Branch{
		OrganizationID: data.OrganizationID,
		Name:           data.Name,
		IsActive:       data.IsActive,
	}

	result := br.db.Create(newBranch)

	return result.Error
}

func (br branchRepositry) Update(id int, data dto.BranchUpdateDTO) error {
	branch := model.Branch{}
	result := br.db.First(branch, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = br.db.Updates(data).Update("updated_at", time.Now())
	if result.Error == nil {
		return result.Error
	}

	return nil
}

func (br branchRepositry) Delete(id int) error {
	return br.db.Delete(&model.Branch{}, "id = ?", id).Error
}
