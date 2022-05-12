package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TableCRUDEndpoints(router *gin.Engine, db *gorm.DB, responseFormatter ResponseFormatterFunc) {
	tableRepo := repository.NewTableRepository(db)
	tableService := service.NewTableService(tableRepo)
	tableController := controller.NewTableController(tableService)

	table := router.Group("/table", middleware.AuthMiddleware())
	table.GET("", responseFormatter(tableController.FindAll))
	table.GET("/:id", responseFormatter(tableController.FindByID))
	table.POST("", responseFormatter(tableController.Create))
	table.PUT("/:id", responseFormatter(tableController.Update))
	table.DELETE("/:id", responseFormatter(tableController.Delete))
}
