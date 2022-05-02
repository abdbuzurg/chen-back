package initialMigration

import (
	"chen/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var routesNotAffectedByRoles = []string{
	"/auth/register",
	"/auth/login",
}

func PermissionChecker(r *gin.Engine, db *gorm.DB) error {
	routes := r.Routes()
	permissions := []model.Permission{}
	err := db.Find(&permissions).Error
	if err != nil {
		return err
	}

	if len(permissions) == 0 {
		return permissionInitialization(db, routes)
	} else if !(len(routesNotAffectedByRoles)+len(permissions) == len(routes)) {
		return permissionUpdate(db, routes)
	}

	return nil
}

func permissionInitialization(db *gorm.DB, routes gin.RoutesInfo) error {

	permissions := []model.Permission{}
	for _, route := range routes {
		method := route.Method
		path := route.Path
		isRestricted := false
		for _, restrictedRoute := range routesNotAffectedByRoles {
			if path == restrictedRoute {
				isRestricted = true
				break
			}
		}
		if isRestricted {
			continue
		}

		permissions = append(permissions, model.Permission{
			Method: method,
			Path:   path,
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		})
	}

	return db.CreateInBatches(&permissions, 3).Error
}

func permissionUpdate(db *gorm.DB, routes gin.RoutesInfo) error {
	err := db.Migrator().DropTable(&model.Permission{})
	if err != nil {
		return err
	}

	return permissionInitialization(db, routes)
}
