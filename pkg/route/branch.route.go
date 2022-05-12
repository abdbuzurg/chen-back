package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BranchCRUDEndpoints(router *gin.Engine, db *gorm.DB, responseFormatter ResponseFormatterFunc) {
	branchRepo := repository.NewBranchRepository(db)
	branchService := service.NewBranchService(branchRepo)
	branchController := controller.NewBranchController(branchService)

	branch := router.Group("/branch", middleware.AuthMiddleware())
	branch.GET("", responseFormatter(branchController.FindAll))
	branch.GET("/:id", responseFormatter(branchController.FindByID))
	branch.POST("", responseFormatter(branchController.Create))
	branch.PUT("", responseFormatter(branchController.Update))
	branch.DELETE("", responseFormatter(branchController.Delete))
}
