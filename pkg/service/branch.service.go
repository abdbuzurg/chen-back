package service

import (
	"chen/pkg/dto"
	"chen/pkg/model"
	"chen/pkg/repository"
)

type BranchService interface {
	FindByID(id int) (model.Branch, error)
	FindAll() ([]model.Branch, error)
	Create(data dto.BranchCreate) error
	Update(id int, data dto.BranchUpdate) error
	Delete(id int) error
}

type branchService struct {
	branchRepository repository.BranchRepository
}

func NewBranchService(repo repository.BranchRepository) BranchService {
	return &branchService{
		branchRepository: repo,
	}
}

func (bs branchService) FindByID(id int) (model.Branch, error) {
	return bs.branchRepository.FindByID(id)
}

func (bs branchService) FindAll() ([]model.Branch, error) {
	return bs.branchRepository.FindAll()
}

func (bs branchService) Create(data dto.BranchCreate) error {
	return bs.branchRepository.Create(data)
}

func (bs branchService) Update(id int, data dto.BranchUpdate) error {
	return bs.branchRepository.Update(id, data)
}

func (bs branchService) Delete(id int) error {
	return bs.branchRepository.Delete(id)
}
