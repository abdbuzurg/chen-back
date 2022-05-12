package service

import (
	"chen/pkg/dto"
	"chen/pkg/model"
	"chen/pkg/repository"
)

type OrganizationService interface {
	FindAll() ([]model.Organization, error)
	FindByID(id int) (model.Organization, error)
	Create(data dto.Organization) error
	Update(id int, data dto.Organization) error
	Delete(id int) error
}

type organizationService struct {
	organizationRepo repository.OrganizationRepo
}

func NewOrganizationService(repo repository.OrganizationRepo) OrganizationService {
	return &organizationService{
		organizationRepo: repo,
	}
}

func (os organizationService) FindByID(id int) (model.Organization, error) {
	return os.organizationRepo.FindByID(id)
}

func (os organizationService) FindAll() ([]model.Organization, error) {
	return os.organizationRepo.FindAll()
}

func (os organizationService) Create(data dto.Organization) error {
	return os.organizationRepo.Create(data)
}
func (os organizationService) Update(id int, data dto.Organization) error {
	return os.organizationRepo.Update(id, data)
}
func (os organizationService) Delete(id int) error {
	return os.organizationRepo.Delete(id)
}
