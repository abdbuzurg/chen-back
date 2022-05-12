package repository

import (
	"chen/pkg/dto"
	"chen/pkg/model"
	"time"

	"gorm.io/gorm"
)

type TableRepository interface {
	FindAll() ([]model.Table, error)
	FindByID(id int) (model.Table, error)
	Create(data dto.Table) error
	Update(id int, data dto.Table) error
	Delete(id int) error
}

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{
		db: db,
	}
}

func (tr tableRepository) FindAll() ([]model.Table, error) {
	tables := []model.Table{}
	err := tr.db.Find(&tables).Error

	return tables, err
}

func (tr tableRepository) FindByID(id int) (model.Table, error) {
	table := model.Table{}
	err := tr.db.First(&table, "id = ?", id).Error
	return table, err
}

func (tr tableRepository) Create(data dto.Table) error {
	table := model.Table{
		X:      data.X,
		Y:      data.Y,
		Number: data.Number,
		HallID: data.HallID,
	}
	err := tr.db.Create(&table).Error
	return err
}

func (tr tableRepository) Update(id int, data dto.Table) error {
	table := model.Table{}

	err := tr.db.First(&table, "id = ?", id).Error
	if err != nil {
		return err
	}

	table.X = data.X
	table.Y = data.Y
	table.Number = data.Number
	table.HallID = data.HallID
	table.UpdatedAt = time.Now()

	err = tr.db.Updates(&table).Error

	return err
}

func (tr tableRepository) Delete(id int) error {
	return tr.db.Delete(&model.Table{}, id).Error
}
