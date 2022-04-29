package service

import (
	"chen/model"
	"chen/pkg/dto"
	"chen/pkg/repository"
)

type OrganizationService interface {
	FindAll() ([]model.Organization, error)
	FindById(id int) (model.Organization, error)
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

func (os organizationService) FindAll() ([]model.Organization, error) {
	return os.organizationRepo.FindAll()
}

func (os organizationService) FindById(id int) (model.Organization, error) {
	return os.organizationRepo.FindById(id)
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
