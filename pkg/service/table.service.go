package service

import (
	"chen/pkg/dto"
	"chen/pkg/model"
	"chen/pkg/repository"
)

type TableService interface {
	FindByID(id int) (model.Table, error)
	FindAll() ([]model.Table, error)
	Create(data dto.Table) error
	Update(id int, data dto.Table) error
	Delete(id int) error
}

type tableService struct {
	tableRepository repository.TableRepository
}

func NewTableService(repo repository.TableRepository) TableService {
	return &tableService{
		tableRepository: repo,
	}
}

func (ts tableService) FindAll() ([]model.Table, error) {
	return ts.tableRepository.FindAll()
}

func (ts tableService) FindByID(id int) (model.Table, error) {
	return ts.tableRepository.FindByID(id)
}

func (ts tableService) Create(data dto.Table) error {
	return ts.tableRepository.Create(data)
}

func (ts tableService) Update(id int, data dto.Table) error {
	return ts.tableRepository.Update(id, data)
}

func (ts tableService) Delete(id int) error {
	return ts.tableRepository.Delete(id)
}
