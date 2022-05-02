package initialMigration

import (
	"chen/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

const ORGANIZATION_ID = 1

type Permission struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type role struct {
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	OrganizationID uint         `json:"organization_id"`
	Permissions    []Permission `json:"permissions"`
}

// open file and parses the file
func openFileAndParse(filename string) ([]role, error) {

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("could not open json file")
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	roles := []role{}
	err = json.Unmarshal(byteValue, &roles)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("could not parse json file")
		return nil, err
	}

	return roles, nil
}

//Check if role table has the default roles
func RoleChecker(db *gorm.DB) error {
	roles := []model.Role{}
	err := db.Find(&roles).Error
	if err != nil {
		return err
	}

	defaultRoles, err := openFileAndParse("db/migration/initialMigration/roles_initial_migration.json")
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		return initializeRoles(db, defaultRoles)
	}

	return updateRoles(db, roles, defaultRoles)
}

// initializes roles if they do not exist
func initializeRoles(db *gorm.DB, defaultRoles []role) error {
	for _, role := range defaultRoles {
		err := db.Create(&role).Error
		if err != nil {
			return err
		}

		for _, permission := range role.Permissions {
			if err = db.Model(&role).Association("permissions").Append(&model.Permission{
				Method: permission.Method,
				Path:   permission.Path,
			}); err != nil {
				return err
			}

		}
	}

	return nil
}

// check if default ROLES are there
// if not returs the STATUS (bool) and the missings ROLES ([]Roles)
func missingDefaultRoles(roles []model.Role, defaultRoles []role) (bool, []model.Role) {
	status := false
	result := []model.Role{}

	for _, defaultRole := range defaultRoles {
		exists := false
		for _, role := range roles {
			if role.Title == defaultRole.Title {
				exists = true
				break
			}
		}

		if !exists {
			result = append(result, model.Role{
				Model: gorm.Model{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
				},
				Title:          defaultRole.Title,
				Description:    defaultRole.Description,
				OrganizationID: ORGANIZATION_ID,
			})

			status = true
		}
	}

	return status, result
}

// updates roles if they are not there
func updateRoles(db *gorm.DB, roles []model.Role, defaultRoles []role) error {
	if status, missingRoles := missingDefaultRoles(roles, defaultRoles); status {
		return db.CreateInBatches(&missingRoles, 4).Error
	}

	return nil
}
