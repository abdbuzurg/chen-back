package migration

import (
	"chen/db/migration/initialMigration"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitialMigration(r *gin.Engine, db *gorm.DB) error {
	if err := initialMigration.PermissionChecker(r, db); err != nil {
		log.Fatalln("could not make initial migration/checkup for permission table")
		return err
	}

	if err := initialMigration.RoleChecker(db); err != nil {
		log.Fatalln("could not make initial migration/checkup for role table")
	}

	return nil
}
