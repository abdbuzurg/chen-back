package repository

import (
	"chen/model"
	"chen/pkg/dto"
	"time"

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
	result := hr.db.First(&hall, "id = ?", id)
	return hall, result.Error
}

func (hr hallRepository) Create(data dto.HallCreateDTO) error {
	hall := &model.Hall{
		Name:     data.Name,
		BranchID: data.BranchID,
	}

	result := hr.db.Create(hall)

	return result.Error
}

func (hr hallRepository) Update(id int, data dto.HallUpdateDTO) error {
	hall := model.Hall{}
	err := hr.db.First(&hall, "id = ?", id).Error
	if err != nil {
		return err
	}

	hall.Name = data.Name

	return hr.db.Updates(hall).Update("updated_at", time.Now()).Error
}

func (hr hallRepository) Delete(id int) error {
	hall := model.Hall{}
	result := hr.db.Delete(&hall, id)
	return result.Error
}
