package route

import (
	"chen/pkg/controller"
	"chen/pkg/middleware"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) HallCRUDEndpoints() {
	hallRepo := repository.NewHallRepository(s.DB)
	hallService := service.NewHallService(hallRepo)
	hallController := controller.NewHallController(hallService)

	hall := s.Router.Group("/hall", middleware.AuthMiddleware())
	hall.GET("", hallController.FindAll)
	hall.GET("/:id", hallController.FindById)
	hall.POST("", hallController.Create)
	hall.PUT("/:id", hallController.Update)
	hall.DELETE("/:id", hallController.Delete)
}
