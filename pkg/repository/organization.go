package repository

import (
	"chen/db"
	"chen/model"
	"errors"
	"time"
)

func OrgFindById(id int) (model.Organization, error) {
	db := db.GetSQLiteConnection()

	var org model.Organization
	result := db.First(&org, "id = ?", id)
	if result.Error != nil {
		return model.Organization{}, errors.New("org not found")
	}

	return org, nil
}
func OrgCreate(data model.OrganizationData) error {
	db := db.GetSQLiteConnection()

	newOrg := &model.Organization{
		IsActive: data.IsActive,
		Name:     data.Name,
	}

	result := db.Create(newOrg)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func OrgUpdate(id int, data model.OrganizationData) error {
	db := db.GetSQLiteConnection()

	var org model.Organization
	result := db.First(&org, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Model(org).Updates(data).Update("updated_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func OrgDelete(id int) error {
	db := db.GetSQLiteConnection()
	return db.Delete(&model.Organization{}, id).Error
}
