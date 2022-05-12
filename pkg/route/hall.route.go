package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HallCRUDEndpoints(router *gin.Engine, db *gorm.DB, responseFromatter ResponseFormatterFunc) {
	hallRepo := repository.NewHallRepository(db)
	hallService := service.NewHallService(hallRepo)
	hallController := controller.NewHallController(hallService)

	hall := router.Group("/hall", middleware.AuthMiddleware())
	hall.GET("", responseFromatter(hallController.FindAll))
	hall.GET("/:id", responseFromatter(hallController.FindByID))
	hall.POST("", responseFromatter(hallController.Create))
	hall.PUT("/:id", responseFromatter(hallController.Update))
	hall.DELETE("/:id", responseFromatter(hallController.Delete))
}
