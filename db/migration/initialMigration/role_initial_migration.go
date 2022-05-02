package initialMigration

import (
	"chen/model"
	"time"

	"gorm.io/gorm"
)

var defaultRoles = map[string]string{
	"SuperAdmin": "Full control in the system",
	"Admin":      "Has full access in the organization/branch",
	"Manager":    "Has access to controlling the stuff and more defined by you",
	"Cashier":    "Has access to finishing order and more defined by you",
	"Waiter":     "Has access to only ordering",
}

const organizationID = 1

func RoleChecker(db *gorm.DB) error {
	roles := []model.Role{}
	err := db.Find(&roles).Error
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		return initializeRoles(db)
	}

	return updateRoles(db, roles)
}

func initializeRoles(db *gorm.DB) error {
	roles := []model.Role{}
	for key, value := range defaultRoles {
		roles = append(roles, model.Role{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			Title:          key,
			Description:    value,
			OrganizationID: organizationID,
		})
	}

	return db.CreateInBatches(&roles, 5).Error
}

// check if default ROLES are there
// if not returs the STATUS (bool) and the missings ROLES ([]Roles)
func missingDefaultRoles(roles []model.Role) (bool, []model.Role) {
	status := false
	result := []model.Role{}

	for key, value := range defaultRoles {
		exists := false
		for _, role := range roles {
			if role.Title == key {
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
				Title:          key,
				Description:    value,
				OrganizationID: organizationID,
			})

			status = true
		}
	}

	return status, result
}

func updateRoles(db *gorm.DB, roles []model.Role) error {
	if status, missingRoles := missingDefaultRoles(roles); status {
		return db.CreateInBatches(&missingRoles, 4).Error
	}

	return nil
}

func destributeDefaultPermission() {

}
