package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) TableCRUDEndpoint() {
	tableRepo := repository.NewTableRepository(s.DB)
	tableService := service.NewTableService(tableRepo)
	tableController := controller.NewTableController(tableService)

	table := s.Router.Group("/table", middleware.AuthMiddleware())
	table.GET("/", tableController.Find)
	table.POST("", tableController.Create)
	table.PUT("/:id", tableController.Update)
	table.DELETE("/:id", tableController.Delete)
}
