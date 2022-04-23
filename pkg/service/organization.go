package service

import (
	"chen/model"
	"chen/pkg/repository"
)

func OrgFind(id int) (model.Organization, error) {

	return repository.OrgFindById(id)
}
func OrgCreate(data model.OrganizationData) error {
	return repository.OrgCreate(data)
}
func OrgUpdate(id int, data model.OrganizationData) (model.Organization, error) {
	return repository.OrgUpdate(id, data)
}
func OrgDelete(id int) error {
	return repository.OrgDelete(id)
}
