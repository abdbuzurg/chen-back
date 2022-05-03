package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type BranchService interface {
	Find(id int) ([]model.Branch, error)
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

func (bs branchService) Find(id int) ([]model.Branch, error) {
	if id == 0 {
		return bs.branchRepository.FindAll()
	}

	branch, err := bs.branchRepository.FindById(id)
	return append([]model.Branch{}, branch), err
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
