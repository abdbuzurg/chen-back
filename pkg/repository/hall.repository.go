package repository

import (
	"chen/model"
	"chen/pkg/dto"

	"gorm.io/gorm"
)

type HallRepository interface {
	FindAll() ([]model.Hall, error)
	FindById(id int) (model.Hall, error)
	Create(data dto.HallCreateDTO) error
	Update(id int, data dto.HallUpdateDTO) error
	Delete(id int) error
}

type hallRepository struct {
	db *gorm.DB
}

func NewHallRepository(db *gorm.DB) HallRepository {
	return hallRepository{
		db: db,
	}
}

func (hr hallRepository) FindAll() ([]model.Hall, error) {
	halls := []model.Hall{}
	err := hr.db.Find(&halls).Error
	return halls, err
}

func (hr hallRepository) FindById(id int) (model.Hall, error) {
	hall := model.Hall{}
	err := hr.db.First(&hall, "id = ?", id).Error
	return hall, err
}

func (hr hallRepository) Create(data dto.HallCreateDTO) error {
	return hr.db.Create(&model.Hall{
		Name:     data.Name,
		BranchID: data.BranchID,
	}).Error
}

func (hr hallRepository) Update(id int, data dto.HallUpdateDTO) error {
	hall := model.Hall{}
	err := hr.db.First(&hall, "id = ?", id).Error
	if err != nil {
		return err
	}

	return hr.db.Updates(hall).Updates(data).Error
}

func (hr hallRepository) Delete(id int) error {
	return hr.db.Delete(model.Hall{}, id).Error
}
