package repository

import (
	"chen/pkg/dto"
	"chen/pkg/model"

	"gorm.io/gorm"
)

type HallRepository interface {
	FindAll() ([]model.Hall, error)
	FindByID(id int) (model.Hall, error)
	Create(data dto.HallCreate) error
	Update(id int, data dto.HallUpdate) error
	Delete(id int) error
}

type hallRepository struct {
	db *gorm.DB
}

func NewHallRepository(db *gorm.DB) HallRepository {
	return &hallRepository{
		db: db,
	}
}

func (hr hallRepository) FindAll() ([]model.Hall, error) {
	halls := []model.Hall{}
	err := hr.db.Find(&halls).Error
	return halls, err
}

func (hr hallRepository) FindByID(id int) (model.Hall, error) {
	hall := model.Hall{}
	err := hr.db.First(&hall, "id = ?", id).Error
	return hall, err
}

func (hr hallRepository) Create(data dto.HallCreate) error {
	return hr.db.Create(&model.Hall{
		Name:     data.Name,
		BranchID: data.BranchID,
	}).Error
}

func (hr hallRepository) Update(id int, data dto.HallUpdate) error {
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
