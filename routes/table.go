package routes

import (
	"chen/pkg/controller"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) TableCRUDEndpoint() {
	tableRepo := repository.NewTableRepository(s.DB)
	tableService := service.NewTableService(tableRepo)
	tableController := controller.NewTableController(tableService)

	table := s.Router.Group("/table")
	table.GET("", tableController.FindAll)
	table.GET("/:id", tableController.FindById)
	table.POST("", tableController.Create)
	table.PUT("/:id", tableController.Update)
	table.DELETE("/:id", tableController.Delete)
}
