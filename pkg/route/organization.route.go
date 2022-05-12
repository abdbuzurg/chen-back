package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrganizationCRUDEndpoints(router *gin.Engine, db *gorm.DB, responseFormatter ResponseFormatterFunc) {
	organizationRepo := repository.NewOrganizationRepo(db)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationController := controller.NewOrganizationController(organizationService)

	org := router.Group("/organization", middleware.AuthMiddleware())
	org.GET("", responseFormatter(organizationController.FindAll))
	org.GET("/:id", responseFormatter(organizationController.FindByID))
	org.POST("", responseFormatter(organizationController.Create))
	org.PUT("/:id", responseFormatter(organizationController.Update))
	org.DELETE("/:id", responseFormatter(organizationController.Delete))
}
