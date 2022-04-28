package routes

import (
	"chen/pkg/controller"
	"chen/pkg/repository"
	"chen/pkg/service"
)

func (s *Server) OrganizationCRUDEndpoints() {
	organizationRepo := repository.NewOrganizationRepo(s.DB)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationController := controller.NewOrganizationController(organizationService)

	org := s.Router.Group("/organization")
	org.GET("", organizationController.FindAll)
	org.GET("/:id", organizationController.FindById)
	org.POST("", organizationController.Create)
	org.PUT("/:id", organizationController.Update)
	org.DELETE("/:id", organizationController.Delete)
}
