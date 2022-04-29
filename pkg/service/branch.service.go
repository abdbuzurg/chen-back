package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type BranchService interface {
	FindAll() ([]model.Branch, error)
	FindById(id int) (model.Branch, error)
	Create(data dto.BranchCreateDTO) error
	Update(id int, data dto.BranchUpdateDTO) error
	Delete(id int) error
}

type branchService struct {
	branchRepository repository.BranchRepository
}

func NewBranchService(repo repository.BranchRepository) BranchService {
	return branchService{
		branchRepository: repo,
	}
}

func (bs branchService) FindAll() ([]model.Branch, error) {
	return bs.branchRepository.FindAll()
}

func (bs branchService) FindById(id int) (model.Branch, error) {
	return bs.branchRepository.FindById(id)
}

func (bs branchService) Create(data dto.BranchCreateDTO) error {
	return bs.branchRepository.Create(data)
}

func (bs branchService) Update(id int, data dto.BranchUpdateDTO) error {
	return bs.branchRepository.Update(id, data)
}

func (bs branchService) Delete(id int) error {
	return bs.branchRepository.Delete(id)
}
