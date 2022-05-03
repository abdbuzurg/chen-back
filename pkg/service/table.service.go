package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type TableService interface {
	Find(id int) ([]model.Table, error)
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

func (ts tableService) Find(id int) ([]model.Table, error) {
	if id == 0 {
		return ts.tableRepository.FindAll()
	}

	table, err := ts.tableRepository.FindById(id)
	return append([]model.Table{}, table), err
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
