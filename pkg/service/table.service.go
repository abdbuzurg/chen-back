package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type TableService interface {
	FindAll() ([]model.Table, error)
	FindById(id int) (model.Table, error)
	Create(data dto.TableDTO) error
	Update(id int, data dto.TableDTO) error
	Delete(id int) error
}

type tableService struct {
	tableRepository repository.TableRepository
}

func NewTableService(repo repository.TableRepository) TableService {
	return tableService{
		tableRepository: repo,
	}
}

func (ts tableService) FindAll() ([]model.Table, error) {
	return ts.tableRepository.FindAll()
}

func (ts tableService) FindById(id int) (model.Table, error) {
	return ts.tableRepository.FindById(id)
}

func (ts tableService) Create(data dto.TableDTO) error {
	return ts.tableRepository.Create(data)
}

func (ts tableService) Update(id int, data dto.TableDTO) error {
	return ts.tableRepository.Update(id, data)
}

func (ts tableService) Delete(id int) error {
	return ts.tableRepository.Delete(id)
}
