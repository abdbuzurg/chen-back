package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type OrganizationService interface {
	Find(id int) ([]model.Organization, error)
	Create(data dto.OrganizationDTO) error
	Update(id int, data dto.OrganizationDTO) error
	Delete(id int) error
}

type organizationService struct {
	organizationRepo repository.OrganizationRepo
}

func NewOrganizationService(repo repository.OrganizationRepo) OrganizationService {
	return organizationService{
		organizationRepo: repo,
	}
}

func (os organizationService) Find(id int) ([]model.Organization, error) {
	if id == 0 {
		return os.organizationRepo.FindAll()
	}

	organization, err := os.organizationRepo.FindById(id)
	return append([]model.Organization{}, organization), err
}
func (os organizationService) Create(data dto.OrganizationDTO) error {
	return os.organizationRepo.Create(data)
}
func (os organizationService) Update(id int, data dto.OrganizationDTO) error {
	return os.organizationRepo.Update(id, data)
}
func (os organizationService) Delete(id int) error {
	return os.organizationRepo.Delete(id)
}
